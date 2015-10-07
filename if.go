package main

import "fmt"

func main() {

	name := "avcssc"

	if len(name) > 0 {
		fmt.Println("aaa")
	}

	mapa := make(map[string]string)
	mapa["name"] = "liujiangbei"
  
	if _, ok := mapa["name"]; ok {
		fmt.Println("GOGOGOOG")
	}
}
