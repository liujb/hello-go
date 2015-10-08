package main

import "fmt"

// 函数变参
func funcParams(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	funcParams(1, 2, 3, 4)
}
