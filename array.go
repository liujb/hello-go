package main

import "fmt"

func testArray() {
	array := [...]int{12, 23, 34, 56, 4, 56, 45, 6, 46, 4, 6, 54, 64, 6, 4, 64, 64, 64, 6, 4, 64, 6}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func main() {
	testArray()
}
