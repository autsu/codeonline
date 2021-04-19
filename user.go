package main

import (
	"math/rand"
	"strconv"
	"time"
)

type IP = string

type User struct {
	Filename string
	Code
}

type Code struct {
	Type    string
	Content string
}

func NewUser(codeType, codeContent string) *User {
	rand.Seed(time.Now().UnixNano())
	userId := strconv.Itoa(rand.Intn(10000))

	filename := codeType + userId

	return &User{
		Filename: filename,
		Code: Code{
			Type:    codeType,
			Content: codeContent,
		},
	}
}

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
