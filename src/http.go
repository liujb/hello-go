package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// 创建链接和处理函数
	http.HandleFunc("/", handler)

	// 监听端口
	http.ListenAndServe(":8089", nil)
}
