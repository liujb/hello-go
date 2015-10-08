package main

import "fmt"

func main() {

	// 定义slice
	slice := []int{1, 2, 3, 4, 5, 6}

	// 循环slice
	for k, v := range slice {
		fmt.Println(k, v)
	}

	// 定义并初始化map
	maps := make(map[string]int)

	maps["one"] = 1
	maps["two"] = 2

	for k, v := range maps {
		fmt.Println(k, v)
	}
}
