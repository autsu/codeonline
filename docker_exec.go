package main

import (
	"errors"
	"log"
	"os/exec"
)

// DockerCP Copy local files to the docker container，or docker to local
// usage: docker cp test.go [container ID or Name]:/
// example: docker cp test.go gocode11:/
func DockerCP(user *User) error {
	switch user.Type {
	case TypeGO:
		//  docker cp /Users/abc/GolandProjects/tools/codeonline/code/1234.go gotest:code/
		c := "docker cp " + TempFilePath + user.Filename + " " + ContainerGo + ":" + SourceFilePath
		cmd := exec.Command("bash", "-c", c)
		log.Println("docker cp command: ", cmd)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return errors.New(err.Error() + string(output))
		}
	}
	return nil
}

func DockerRun() {
	exec.Command("docker", "run", "")
}

// DockerExecAndRunCode 进入容器内部并执行源文件
// usage: 先 exec 进入容器，再在容器中执行命令
// docker exec -it gocode11 sh -c " ls -l && go run test.go"
func DockerExecAndRunCode(user *User) []byte {
	switch user.Type {
	case TypeGO:
		// docker exec -i gotest sh -c "go run 1420.go"
		cmd := "docker exec -i " + ContainerGo + ` sh -c "cd ../code && ` + "go run " + user.Filename + `"`
		log.Println("docker exec and run command: ", cmd)
		com := exec.Command("bash", "-c", cmd)
		output, _ := com.CombinedOutput()
		return output
	}
	return []byte("未知/不支持的语言")
}

// DockerExecAndRemoveFile 删除容器内部的代码文件
func DockerExecAndRemoveFile(user *User) error {
	switch user.Type {
	case TypeGO:
		// docker exec and rm command:  docker exec -i gotest sh -c "cd ../code && rm 214.go"
		cmd := "docker exec -i " + ContainerGo + ` sh -c "cd ../code && ` + "rm " + user.Filename + `"`
		log.Println("docker exec and rm command: ", cmd)
		com := exec.Command("bash", "-c", cmd)
		output, err := com.CombinedOutput()
		if err != nil {
			return errors.New(err.Error() + string(output))
		}
	}
	return nil
}
