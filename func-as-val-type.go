package main

import "fmt"

// 定义函数变量
type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 54, 34, 67, 8}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd)
	even := filter(slice, isEven)

	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)
}
