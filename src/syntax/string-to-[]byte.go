package main

import "fmt"

func main() {
	str := "hhiello, world"
	fmt.Println("len", len(str))
	// fmt.Println("len", utf8.RuneCountInString(str))

	// 将str转换为byte切片
	bytes := []byte(str)
	// fmt.Println("bytes len", len(bytes))

	for _, v := range bytes {
		// fmt.Println(reflect.TypeOf(v))
		fmt.Println(string(v))
	}
}
