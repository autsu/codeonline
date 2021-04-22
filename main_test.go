package main

import (
	"encoding/json"
	"log"
	"testing"
)

func TestRecJson(t *testing.T) {
	js := `
{
"type": "go",
"content": "package main\n import (\n \"fmt\" \n) \n func main() {\n fmt.Println(\"test\") \n } \n "
}
`
	// 死循环
	_ = `
{
"type": "go",
"content": "package main\nimport()\nfunc main(){\nfor{}\n}"
}
`
	//log.Println(js)
	var code Code
	if err := json.Unmarshal([]byte(js), &code); err != nil {
		log.Fatal(err)
	}
	log.Println(code.Type)
	log.Println(code.Content)
}
