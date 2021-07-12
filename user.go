package main

import (
	"encoding/json"
	"io"
	"log"
)

type User struct {
	// 标识一个文件名，当在 docker 中创建代码文件时会使用该字段，此外，
	// 在 docker 中 build 以及 run 都会使用该字段，这里写死为 test，
	// 之后再根据代码类型加上对应的后缀名即可
	Filename string

	// ContainerName 是容器名，在 DockerRun() 中为该字段赋值，
	// 当执行 docker run 创建容器时，会使用该名称来命名
	ContainerName string

	// 用户的执行代码
	Code *Code
}

type Code struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func NewCode(reqBody io.ReadCloser) (*Code, error) {
	body, err := io.ReadAll(reqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var code Code
	if err := json.Unmarshal(body, &code); err != nil {
		log.Println(err)
		return nil, err
	}

	return &code, nil
}

func NewUser(code *Code) *User {
	// 文件名写死为 test
	filename := "test"

	switch code.Type {
	case TypeGO:
		filename += SuffixGo
	case TypeJava:
		// 这里写死，因为 java 编译需要保证类名和文件名相同
		filename += SuffixJava
	case TypeC:
		filename += SuffixC
	case TypeCPP:
		filename += SuffixCPP
	}
	return &User{
		Code:     code,
		Filename: filename,
	}
}

