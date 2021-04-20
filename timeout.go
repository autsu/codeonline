package main

import (
	"errors"
	"log"
	"time"
)

// 超时时间
const timeout = time.Second * 6

type RunFunc = func(*User) []byte

func Timeout(user *User, f RunFunc) (result []byte, err error) {
	done := make(chan struct{})

	go func() {
		// 如果 f() 是死循环，会导致 goroutine 泄漏
		res := f(user)
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
