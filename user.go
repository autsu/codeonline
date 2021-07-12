package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"time"
)

//type IP = string

type User struct {
	// 用 ip 可以标识用户，Id 字段不再需要了
	//Id string

	// Filename 随机文件名保证一个用户的容器下可以存放多个代码文件
	// TODO 在新的 timeout 解决策略下（具体参加 timeout.go 注释），
	// 文件名似乎没有什么意义了，这个字段在未来可能会删除
	//
	// ========  分割线 ========
	//
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
	rand.Seed(time.Now().UnixNano())
	// 文件名不需要随机生成了
	//userId := strconv.Itoa(rand.Intn(10000))

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
		//Id:       userId,
		Code:     code,
		Filename: filename,
	}
}

// Users 暂时没什么用
//type Users map[IP]*User
//
//func NewUsers() Users {
//	return make(Users)
//}
//
//func (u Users) Add(ip IP, user *User) {
//	u[ip] = user
//}
//
//func (u Users) Delete(ip string) {
//	delete(u, ip)
//}

//func (u Users) RegisterUser(addr string, reqBody io.ReadCloser) error {
//	body, err := io.ReadAll(reqBody)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	var code Code
//	if err := json.Unmarshal(body, &code); err != nil {
//		log.Println(err)
//		return err
//	}
//	us := NewUser(code)
//	u.Add(addr, us)
//	return nil
//}
