package main

import "fmt"

// func 返回多个值
func multi() (string, int) {
	return "nihao", 10
}

func main() {
	x, y := multi()
	fmt.Println(x)
	fmt.Println(y)
}
