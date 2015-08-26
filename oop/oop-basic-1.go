package main

import "fmt"

// 声明对象
type Rectangle struct {
	width  float64
	height float64
}

// 声明area方法，接收Ractangle类型的对象作为参数
func area(r Rectangle) float64 {
	return r.width * r.height
}

// 入口function
func main() {
	r1 := Rectangle{10, 10}
	fmt.Println("area of r1 is", area(r1))
}
