package main

import "fmt"

func main() {

	// 为map分配内存
	mapSS := make(map[string]string)

	mapSS["one"] = "string"
	if _, ok := mapSS["one"]; ok {
		fmt.Println(ok)
	}

}
