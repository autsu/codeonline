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
	Id            string
	Filename      string
	ContainerName string
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
	}
	return &User{
		Id:       userId,
		Code:     code,
		Filename: filename,
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
