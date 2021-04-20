package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var us1 = NewUsers()

func TestNewUser(t *testing.T) {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		us1.RegisterUser(addr, r.Body)
		// Output:
		// &{Filename:9682.go Code:{Type:go Content:package main}}
		for ip, user := range us1 {
			fmt.Println(ip)
			fmt.Printf("%+v\n", user)
		}
	})
	if err := http.ListenAndServe(":7775", nil); err != nil {
		fmt.Println(err)
		return
	}
}

// 随机重复率测试
func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	m := make(map[int]struct{})
	count := 0

	for i := 0; i < 1000; i++ {
		intn := rand.Intn(1000000)
		if _, ok := m[intn]; ok {
			fmt.Println("========== 重复 ==========", intn)
			count++
		}
		m[intn] = struct{}{}
	}
	fmt.Println("重复次数：", count)
}

func TestRecJson(t *testing.T) {
	js := `
{
	"type": "go",
"content": "package main\n import (\n \"fmt\" \n) \n func main() {\n fmt.Println(\"test\") \n } \n "
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
