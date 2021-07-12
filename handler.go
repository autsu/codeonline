package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//var us = NewUsers()

var limiter = NewIPRateLimiter(1, 5)

type Hand func(w http.ResponseWriter, r *http.Request)

func LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr)
		l := limiter.Limiter(r.RemoteAddr)
		if !l.Allow() {
			log.Println("limiter not allow")
			http.Error(w, "操作过于频繁，请稍后再试", http.StatusTooManyRequests)
			return
		}
		log.Println("limiter allow")
		next.ServeHTTP(w, r)
	})
}

func CodeHandler() Hand {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
		// header 的类型
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")


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
	//if err := us.RegisterUser(addr, r.Body); err != nil {
	//	return err
	//}
	//user := us[addr]

	// 1. 注册用户
	code, err := NewCode(r.Body)
	if err != nil {
		return err
	}
	user := NewUser(code)

	// 2. 将 ip 添加到集合中
	limiter.AddIP(addr)

	// 8. 删除容器
	defer func() {
		if err := DockerRM(user); err != nil {
			log.Println(err)
			return
		}
	}()

	// 3. 创建容器
	if err := DockerRun(user); err != nil {
		return err
	}

	// 4. 在容器中创建 /code/
	if err := DockerExecAndCreateDir(user); err != nil {
		return err
	}

	// 5. 将用户输入的代码写入到临时文件，再将临时文件 cp 到容器中，再 rm 临时文件
	if err := WCPRM(user); err != nil {
		log.Println(err)
		return err
	}

	// 6. DockerExecAndRunCode 进入容器内部并执行用户的代码文件，包裹 Timeout 进行超时处理
	res, err := Timeout(user, DockerExecAndRunCode)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("run code result: ", string(res))

	// 7. 将执行结果返回给用户
	_, err = io.Copy(w, bytes.NewReader(res))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// WCPRM 先将输入的代码写入到本地的文件（w），再 docker cp 到容器中（cp），之后 rm 本地文件（rm）
func WCPRM(user *User) (err error) {
	log.Println("user code: ", user.Code.Content)
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
