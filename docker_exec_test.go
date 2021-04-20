package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestExecAndRunCode(t *testing.T) {
	// error: the input device is not a tty
	// 解决：docker exec -i（去掉 -t）
	// 因为 -t 是分配一个伪终端
	fliename := "test.go"
	c := `docker exec -i gotest sh -c "cd ` + SourceFilePath + `&& go run ` + fliename + `"`
	log.Println(c)
	cmd := exec.Command("bash", "-c", c)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}

// error: 容器内没有 touch 命令
func TestExecAndTouch(t *testing.T) {
	c := `docker exec -i gotest sh -c "cd code" && "touch sou.go"`
	cmd := exec.Command("bash", "-c", c)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}
