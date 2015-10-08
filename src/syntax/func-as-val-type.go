package main

import "fmt"

// 定义函数变量
type testInt func(int) bool

// 判断偶数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

// 判断奇数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 过滤slice，传入不同的handler
func filter(slice []int, fn testInt) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
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
