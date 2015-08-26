package main

import "fmt"

// 声明Rectangle类型
type Rectangle struct {
	width  float32
	heigth float32
}

// 为Rectangle对象
func (this Rectangle) area() float32 {
	return this.heigth * this.width
}

//func (i int) toString() string {
//	return strconv.Itoa(i)
//}

func main() {

	r1 := Rectangle{width: 10, heigth: 10}
	fmt.Println("r1 area is:", r1.area())

}
