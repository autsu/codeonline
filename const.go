package main

// 代码类型
const (
	TypeGO   = "go"
	TypeJava = "java"
)

const (
	// SourceFilePath 所有源代码文件都保存在对应容器下的 code/ 目录中
	SourceFilePath = "/code/"
	// TempFilePath 用于存放零时文件，在 cp 到容器以后会被 rm
	//TempFilePath = "/root/"
	TempFilePath = "/Users/zz/GolandProjects/tools/codeonline/code/"
)

// 容器名
const (
	//ContainerGo   = "gocode11"
	ContainerGo   = "gotest" // 开发下的测试
	ContainerJava = "javacode11"
)

// 后缀名
const (
	SuffixGo   = ".go"
	SuffixJava = ".java"
)
