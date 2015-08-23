package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	for k, v := range slice {
		fmt.Println(k, v)
	}

	maps := make(map[string]int)
	maps["one"] = 1
	maps["two"] = 2

	for k, v := range maps {
		fmt.Println(k, v)
	}
}
