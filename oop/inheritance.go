package main

import "fmt"

// 定义Human对象
type Human struct {
	name  string
	age   int
	phone string
}

// 定义Student对象
type Student struct {
	Human  // 匿名字段实现继承
	school string
}

//定义Employee对象
type Employee struct {
	Human
	company string
}

// 为Human对象定义SayHi方法
func (h *Human) SayHi() {
	fmt.Printf("I'm %s, my phone is %s\n", h.name, h.phone)
}

// Employee对象重写SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("I'm employee, I am in %s \n", e.company)
}

func main() {
	student := Student{Human{"liujiangbei", 25, "13269167341"}, "北京大学"}
	workder := Employee{Human{"liujiangbei", 25, "13269167341"}, "滴滴打车"}

	student.SayHi()
	workder.SayHi()
}
