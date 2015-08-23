package main

import "fmt"

func main() {
	// var numbers map[string]int
	// numbers = make(map[string]int, 10)

	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println("第三个数字是:", numbers["three"])
	fmt.Println("len is ", len(numbers))
}
