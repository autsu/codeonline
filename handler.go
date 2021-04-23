package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var us = NewUsers()

type Hand func(w http.ResponseWriter, r *http.Request)

func GoHandler() Hand {
	return func(w http.ResponseWriter, r *http.Request) {
		//io.Copy(w, strings.NewReader("connect success\n"))
		log.Printf("[go] a client[%v][%v] in...\n", r.RemoteAddr, r.Method)

		if err := start(w, r); err != nil {
			log.Println("start error: ", err)
			io.Copy(w, strings.NewReader(err.Error()))
			return
		}
	}
}

// read, write and run
func start(w http.ResponseWriter, r *http.Request) error {
	addr := r.RemoteAddr
	// 1. 注册用户
	if err := us.RegisterUser(addr, r.Body); err != nil {
		return err
	}
	user := us[addr]

	// 7. 删除容器
	defer func() {
		if err := DockerRM(user); err != nil {
			log.Println(err)
			return
		}
	}()

	// 2. 创建容器
	if err := DockerRun(user); err != nil {
		return err
	}

	// 3. 在容器中创建 /code/
	if err := DockerExecAndCreateDir(user); err != nil {
		return err
	}

	// 4. 将用户输入的代码写入到临时文件，再将临时文件 cp 到容器中，再 rm 临时文件
	if err := WCPRM(user); err != nil {
		log.Println(err)
		return err
	}

	// 5. DockerExecAndRunCode 进入容器内部并执行用户的代码文件，包裹 Timeout 进行超时处理
	res, err := Timeout(user, DockerExecAndRunCode)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("run code result: ", string(res))

	// 6. 将执行结果返回给用户
	_, err = io.Copy(w, bytes.NewReader(res))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// WCPRM 先将输入的代码写入到本地的文件（w），再 docker cp 到容器中（cp），之后 rm 本地文件（rm）
func WCPRM(user *User) (err error) {
	// 3. 删除临时文件
	// 不论上面的步骤是否执行成功，都删除临时文件
	defer func() {
		// 3. rm 本地临时文件
		if errs := os.Remove(TempFilePath + user.Filename); err != nil {
			err = errs
			return
		}
		log.Println("rm temp file")
	}()
	// 1. 将输入的代码写入到本地的文件
	if err := WriteToTempFile(user); err != nil {
		log.Println(err)
		return err
	}

	// 2. cp 到 docker 容器中
	if err := DockerCP(user); err != nil {
		return err
	}

	return nil
}
