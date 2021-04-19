package main

import (
	"context"
	"errors"
	"log"
	"os/exec"
)

func BuildGo() error {
	cmd := exec.Command("go", "build", Go)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("build file error: " + err.Error())
	}
	log.Printf("output: \n%v\n", string(output))
	return nil
}

func RunGo() []byte {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", Go)
	output, _ := cmd.CombinedOutput()
	log.Println(string(output))
	log.Printf("output: \n%v\n", string(output))
	return output
}
