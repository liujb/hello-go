package main

import (
	"fmt"
)

var dst []byte

func main() {
	slice := []byte{'h', 'h', 'n', 'm'}

	for k, v := range slice {
		fmt.Println(k, v)
	}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := append(slice, 'a')

	fmt.Println(len(slice))
	fmt.Println(len(slice2))

	num := copy(dst, slice)
	fmt.Println(len(dst))
	fmt.Println(num)
}
