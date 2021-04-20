package main

import (
	"bytes"
	"io"
	"os"
)

var (
	// Go = /root/code_online/code/xxx.go
	// Go = /Users/zz/GolandProjects/tools/codeonline/code/xxx.go
	Go = "/code/xxx.go"
)

func WriteToTempFile(user *User) error {
	// TODO 文件加随机数，区分不同用户
	file, err := os.OpenFile(TempFilePath+user.Filename,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(user.Content)
	if err != nil {
		return err
	}
	return nil
}

func ReadBody(body io.Reader) (res []byte, err error) {
	buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	io.Copy(buffer, body)
	temp := buffer.Bytes()
	length := len(temp)
	if cap(temp) > (length + length/10) {
		res = make([]byte, length)
		copy(res, temp)
	} else {
		res = temp
	}
	return
}
