package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, s := 10, "10"

	fmt.Println(strconv.Itoa(i))
	fmt.Println(strconv.Atoi(s))

	fmt.Println(i)
	fmt.Println(s)
}
