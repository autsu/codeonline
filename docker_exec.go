package main

import "os/exec"

// DockerCP Copy local files to the docker container，or docker to local
// usage: docker cp test.go [container ID or Name]:/
// example: docker cp test.go gocode11:/
func DockerCP(codeType string) {
	switch codeType {
	case TypeGO:
		exec.Command("docker", "cp", ContainerGo+":", SourceFilePath)
	}
}

func DockerRun() {
	exec.Command("docker", "run", "")
}

// DockerExec
// usage: 先 exec 进入容器，再在容器中执行命令
// docker exec -it gocode11 sh -c " ls -l && go run test.go"
func DockerExec(codeType, codeFileName string) {
	switch codeType {
	case TypeGO:
		// TODO
		cmd := `sh -c "go run " +  `
		exec.Command("docker", "exec", "-it", ContainerGo, cmd)
	}
}

func D() {

}
