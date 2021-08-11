package main

// source code type
const (
	TypeGO   = "Go"
	TypeJava = "Java"
	TypeC    = "C"
	TypeCPP  = "C++"
)

const (
	// SourceFilePath all source code files are stored in the "/code/" under the corresponding container
	SourceFilePath = "/code/"

	// TempFilePath store the temp file, when copy to container finish, this file will be remove

	TempFilePath = "/root/"	// 部署路径
	//TempFilePath = "/Users/zz/" // 本机测试
)

// 镜像名
const (
	ImageGo   = "golang"
	ImageJava = "java"
	ImageC    = "gcc"
	ImageCPP  = "gcc"
)

// 后缀名
const (
	SuffixGo   = ".go"
	SuffixJava = ".java"
	SuffixC    = ".c"
	SuffixCPP  = ".cc"
)
