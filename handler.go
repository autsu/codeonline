package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Hand func(w http.ResponseWriter, r *http.Request)

func GoWithText() Hand {
	return func(w http.ResponseWriter, r *http.Request) {
		//io.Copy(w, strings.NewReader("connect success\n"))
		log.Printf("[go] a client[%v][%v] in...\n", r.RemoteAddr, r.Method)
		log.Printf("client send body: ")
		//io.Copy(os.Stdout, r.Body)
		fmt.Println()
		start(w, r)
	}
}

// read, write and run
func start(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println("request body: ", body)

	if err = WriteToFile(string(body)); err != nil {
		log.Println(err)
		return
	}

	res, err := Timeout(RunGo)
	//res, err := util.RunGo()
	if err != nil {
		log.Println(err)
		io.Copy(w, strings.NewReader(err.Error()))
		return
	}

	_, err = io.Copy(w, bytes.NewReader(res))
	if err != nil {
		log.Println(err)
		return
	}
}
