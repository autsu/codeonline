package main

import (
	"golang.org/x/time/rate"
	"sync"
)

type IP = string
type IPS map[IP]*rate.Limiter

type IPRateLimiter struct {
	ips IPS          // 所有的 ip，每个 ip 有一个对应的限流器
	mu  sync.RWMutex // 读写锁
	r   rate.Limit   // 每秒 r 个令牌
	b   int          // 最大突发事件
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(IPS),
		mu:  sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

func (i *IPRateLimiter) AddIP(ip IP) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	l := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = l

	return l
}

func (i *IPRateLimiter) Limiter(ip IP) *rate.Limiter {
	i.mu.Lock()
	// 使用 defer 会导致死锁，参考下面的注释
	//defer i.mu.Unlock()

	l, ok := i.ips[ip]
	if !ok {
		// 这里必须 Unlock，因为到这里会调用 AddIP()，而 AddIP() 的第一行就会
		// Lock()，如果这里不 Unlock，会导致 AddIP() 对一个已处于加锁状态的锁
		// 再次加锁，这样会产生死锁
		//
		// 如果使用 defer 释放锁，则会在该函数执行完成后才释放锁，但是这里调用了另一个
		// 函数 AddIP()，进入 AddIP() 会导致死锁
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()
	return l
}
