package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

// 先将输入的代码写入到本地的文件（w），在 docker cp 到容器中（cp），之后 rm 本地文件（rm）
func TestWCPRM(t *testing.T) {
	filename := "/Users/zz/GolandProjects/tools/codeonline/code/noname.go"
	// 1. write code to local file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal("open file error: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(`
package main

import (
    "fmt"
)

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello, World!")
    }
}
`)
	if err != nil {
		log.Fatal(err)
	}

	// 2. docke cp
	c := "docker cp " + filename + " gotest:go/code/"
	cmd := exec.Command("bash", "-c", c)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("docker cp error: ", err)
	}
	fmt.Println(string(output))

	// 3. remove local file
	if err := os.Remove(filename); err != nil {
		log.Fatal(err)
	}
}
