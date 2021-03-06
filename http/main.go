package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Http1 Hello there!\n")
	response, err := http.Get("http://127.0.0.1:3002")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	addr := "0.0.0.0:3001"
	log.Printf("http-server started [%s]", addr)
	http.HandleFunc("/", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(addr, nil))
}
