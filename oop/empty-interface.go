package main

import "fmt"

func main() {
	var a interface{}
	var i int = 5
	var s string = "hello wolrd"

	a = i
	fmt.Println(a)

	a = s
	fmt.Println(a)
}
