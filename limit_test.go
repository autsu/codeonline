package main

import (
	"log"
	"net/http"
	"testing"
)

func TestLimitMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":8081", LimitMiddleware(mux)); err != nil {
		log.Fatalln("start http server error: ", err)
		return
	}
}