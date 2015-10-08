package main

import "fmt"

func testArray() {

	// 数组长度由元素个数推导而来
	array := [...]int{12, 23, 34, 56, 4, 56, 45, 6, 46, 4, 6, 54, 64, 6, 4, 64, 64, 64, 6, 4, 64, 6, 12}
	for i := 0; i < len(array); i++ {
		fmt.Println("number", array[i])
	}
}

func main() {
	testArray()
}
