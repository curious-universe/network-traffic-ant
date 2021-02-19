package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Http2 Hello there!\n")
}

func main() {
	addr := "0.0.0.0:3002"
	log.Printf("http-server started [%s]", addr)
	http.HandleFunc("/", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(addr, nil))
}
