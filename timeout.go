package main

import (
	"errors"
	"log"
	"time"
)

// 超时时间
const timeout = time.Second * 6

type RunFunc = func(*User) ([]byte, error)

func Timeout(user *User, f RunFunc) (result []byte, err error) {
	// 这里需要将缓存设置为 1，如果为无缓存 chan，会导致当 Timeout 函数退出时，
	// 所有请求阻塞在 done <- struct{}{}（因为此时 done 已经不存在了，没有接受者），
	// 从而导致 goroutine 无法释放
	done := make(chan struct{}, 1)

	go func() {
		// 如果 f() 是死循环，会导致 goroutine 泄漏
		// 还没有找到好的方法解决该问题
		//
		// 解决方法：每次运行代码都会创建一个容器，执行完成或者超时
		// 则将该容器强制删除即可，进程结束泄露的 goroutine 自然
		// 也一起结束了
		res, _ := f(user)
		log.Println("执行完成")
		done <- struct{}{}
		result = res
	}()

	// 这里能保证 Timeout 函数退出，但是内部的 goroutine 可能会泄漏
	select {
	case <-time.After(timeout):
		log.Println("函数执行超时")
		return nil, errors.New("函数执行超时")
	case <-done:
		return
	}
}
