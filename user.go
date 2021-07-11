package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type IP = string

type User struct {
	Id string

	// Filename 随机文件名保证一个用户的容器下可以存放多个代码文件
	// TODO 在新的 timeout 解决策略下（具体参加 timeout.go 注释），
	// 文件名似乎没有什么意义了，这个字段在未来可能会删除
	Filename string

	// ContainerName 是容器名，在 DockerRun 中为该字段赋值，
	// 当执行 docker run 创建容器时，会使用该名称来命名
	ContainerName string

	// 用户的执行代码
	Code
}

type Code struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func NewUser(code Code) *User {
	rand.Seed(time.Now().UnixNano())
	userId := strconv.Itoa(rand.Intn(10000))
	filename := userId
	switch code.Type {
	case TypeGO:
		filename += SuffixGo
	case TypeJava:
		// 这里暂时写死，因为 java 编译需要保证类名和文件名相同
		filename = "test" + SuffixJava
	}
	return &User{
		Id:       userId,
		Code:     code,
		Filename: filename,
	}
}

// Users 暂时没什么用
type Users map[IP]*User

func NewUsers() Users {
	return make(Users)
}

func (u Users) Add(ip IP, user *User) {
	u[ip] = user
}

func (u Users) Delete(ip string) {
	delete(u, ip)
}

func (u Users) RegisterUser(addr string, reqBody io.ReadCloser) error {
	body, err := io.ReadAll(reqBody)
	if err != nil {
		log.Println(err)
		return err
	}
	var code Code
	if err := json.Unmarshal(body, &code); err != nil {
		log.Println(err)
		return err
	}
	us := NewUser(code)
	u.Add(addr, us)
	return nil
}
