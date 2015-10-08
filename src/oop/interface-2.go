package main

import "fmt"

type Human struct {
	name  string
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

// 这个地方不能使用 *Human
func (h Human) SayHi() {
	fmt.Println("I am ", h.name, "My phone is", h.phone)
}

// 这个地方不能使用 *Human
func (h Human) Sing(lynics string) {
	fmt.Printf("歌词是：%s\n", lynics)
}

type Men interface {
	SayHi()
	Sing(lynics string)
}

func main() {
	mike := Student{Human{"liujiangbei", "13269167341"}, "北京大学"}
	tom := Employee{Human{"liujb", "18632222222"}, "滴滴打车"}

	var men Men

	men = mike
	men.SayHi()
	men.Sing("难以忘记初次见你，一双迷人的眼睛。")

	men = tom
	men.SayHi()
	men.Sing("你问我爱你有多深")

	slice := make([]Men, 2)
	slice[0], slice[1] = mike, tom
	for _, v := range slice {
		v.SayHi()
	}
}
