package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// 超时时间
const timeout = time.Second * 6

type RunFunc = func() []byte

func Timeout(f RunFunc) (result []byte, err error) {
	done := make(chan struct{})

	go func() {
		// 如果 f() 是死循环，会导致 goroutine 泄漏
		res := f()
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

// Timeout1 插入代码解决
// 废弃
func Timeout1(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {

	}
	code, err := io.ReadAll(file)
	if err != nil {

	}
	s := string(code)
	// 获取
	index := strings.Index(s, "func main")
	for i := index + len("func main"); ; i++ {
		if s[i] == '{' {
			index = i + 1
			break
		}
	}

	cc := "ctx, cancel := context.WithTimeout(context.Background);defer cancel();"
	for i := 0; i < len(cc); i++ {
		file.WriteAt([]byte(" "), int64(index))
		index += 1
	}
	//n, _ :=
	//index += 6 * n

	//file.WriteAt([]byte(cc), int64(index))
	//file.WriteAt([]byte("i := 10"), int64(index+2))
	fmt.Println(index)
}
