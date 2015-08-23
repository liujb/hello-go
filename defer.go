package main

import "fmt"

// defer会按照逆序执行
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
