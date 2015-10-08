package main

import "fmt"

var sum int = 0

// func add
func add() {
	for i := 0; i < 100; i++ {
		sum += i
	}
}

// func main
func main() {
	add()
	fmt.Println(sum)
}
