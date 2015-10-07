package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "Name: " + p.name + " Age: " + strconv.Itoa(p.age)
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"liujiangbei", 25}

	for k, v := range list {

		// comma-ok的使用
		if value, ok := v.(int); ok {
			fmt.Printf("list[%d] is an int, value is %d\n", k, value)
		} else if value, ok := v.(string); ok {
			fmt.Printf("list[%d] is string, value is %s\n", k, valßue)
		}

		// element.(type)不能在switch外的任何逻辑中使用
		switch value := v.(type) {
		case int:
			fmt.Printf("int value is %d\n", value)
		case string:
			fmt.Printf("string value is %s\n", value)
		case Person:
			fmt.Printf("Person value is %s\n", value)
		default:
			fmt.Printf("dddd")
		}
	}
}
