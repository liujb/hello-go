package main

import "fmt"

func main() {
	mapSS := make(map[string]string)
	mapSS["one"] = "string"
	if _, ok := mapSS["one"]; ok {
		fmt.Println(ok)
	}
}
