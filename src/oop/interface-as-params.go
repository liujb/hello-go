package main

import "fmt"

type Human struct {
	name  string
	phone string
}

// Human对象实现了String()方法
// 所以可以被fmt.Println调用
func (h Human) String() string {
	return "Name:" + h.name + " Phone:" + h.phone
}

func main() {
	bob := Human{"liujiangbei", "13269167341"}
	fmt.Println(bob)
}
