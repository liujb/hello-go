package main

import "fmt"

func main() {
	str := "hello, world"

	// 将str转换为byte切片
	bytes := []byte(str)
	fmt.Println(len(bytes))

	for _, v := range bytes {
		fmt.Println(v)
	}
}
