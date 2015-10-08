package main

import "fmt"
import "reflect"

func main() {
	data := []string{"one", "two", "three"}
	test(data)
	moredata := []int{1, 2, 3}
	test(moredata)
}

func test(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}
