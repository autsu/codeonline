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
		// docker cp /Users/abc/GolandProjects/tools/codeonline/code/1234.go gotest:code/
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

func DockerRun(user *User) {
	var c string
	switch user.Type {
	case TypeGO:
		// docker run -i --name gocode11 golang
		c = "docker run -i --name " + ContainerGo + " " + ImageGo
	}
	log.Println("docker run: ", c)
	cmd := exec.Command("bash", "-c", c)
	output, _ := cmd.CombinedOutput()
	log.Println(string(output))
}

// DockerExecAndRunCode exec container and execution this user's code, and return the result
// Example: docker exec -it gocode11 sh -c " ls -l && go run test.go"
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
	return []byte("unknown/not support language")
}

// DockerExecAndRemoveFile exec container, and then remove this user's code file
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

// DockerExecAndCreateDir exec container and create dir "code" in root path "/"
// the user code is store in here
func DockerExecAndCreateDir(user *User) {
	var c string
	switch user.Type {
	case TypeGO:
		c = "docker exec -i" + ContainerGo + ` sh -c "cd .. && mkdir code"`
	}
	log.Println("docker exec and create: ", c)
	cmd := exec.Command("bash", "-c", c)
	output, _ := cmd.CombinedOutput()
	log.Println(string(output))
}

// DockerRM	remove the container according to the code type
func DockerRM(user *User) {
	var c string
	switch user.Type {
	case TypeGO:
		c = "docker rm -f " + ContainerGo
	}
	log.Println("docker rm: ", c)
	cmd := exec.Command("bash", "-c", c)
	output, _ := cmd.CombinedOutput()
	log.Println(string(output))
}
