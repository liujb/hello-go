package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, s, f := 65, "10", 1.234

	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(f)

	// 将数字类型转换为字符串
	// 如果数字65，那么转换字符串即为A
	fmt.Println(string(i))
	fmt.Println(strconv.Itoa(i))
	fmt.Println(strconv.Atoi(s))
}
