package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExec(t *testing.T) {
	output, err := exec.Command("sh", "/Users/zz/GolandProjects/tools/codeonline/sh/docker_exec.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}
