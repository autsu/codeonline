package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// 20911  xxx          97.7      00:31.47 5/1    0     14     872K   0B     632K   20896 1
// 20896  ___go_build_ 94.8      00:30.92 11/1   0     21     911M+  0B     386M   20896 9328

// build: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main *.go
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", CodeHandler())

	// pprof
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// 用 LimitMiddleware() 包裹 mux 实现限流
	if err := http.ListenAndServe(":10001", LimitMiddleware(mux)); err != nil {
		log.Println("start server error: ", err)
		return
	}
}
