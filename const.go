package main

// source code type
const (
	TypeGO   = "go"
	TypeJava = "java"
)

const (
	// SourceFilePath all source code files are stored in the "/code/" under the corresponding container
	SourceFilePath = "/code/"

	// TempFilePath store the temp file, when copy to container finish, this file will be remove
	TempFilePath = "/root/"
	//TempFilePath = "/Users/xx/GolandProjects/tools/codeonline/"
)

// 容器名
const (
	ContainerGo = "gocode11"
	// ContainerGo   = "gotest" // development environment test
	ContainerJava = "javacode11"
)

// 镜像名
const (
	ImageGo   = "golang"
	ImageJava = "java"
)

// 后缀名
const (
	SuffixGo   = ".go"
	SuffixJava = ".java"
)
