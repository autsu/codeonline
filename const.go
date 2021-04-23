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
	//TempFilePath = "/Users/zz/GolandProjects/tools/codeonline/"
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
