package main

import (
	"fmt"
)

func main() {
	slice := []string{"abc", "def"}
	slice = append(slice, "sdfs", "sdfssfds")

	for k, v := range slice {
		fmt.Println(k, v)
	}
}
