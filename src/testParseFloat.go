package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100cd"
	f, _ := strconv.ParseInt(str, 16, 0)
	fmt.Println(f)

	fmt.Printf("%s\n", string(str[len(str)-1]))

	a := 2.3222222222
	fmt.Printf("%2f", a)
}
