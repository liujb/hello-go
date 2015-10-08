package main

import "fmt"

func main() {
	x := 3
	y := 22

	fmt.Println(x)
	fmt.Println(y)

	pi := &x
	fmt.Println(pi)
	fmt.Println(*pi)

	x++
	fmt.Println(x)
	fmt.Println(*pi)

	*pi++
	fmt.Println(x)
	fmt.Println(*pi)

	pi = &y
	fmt.Println(*pi)

	*pi++
	fmt.Println(pi)
	fmt.Println(*pi)
	fmt.Println(y)

}
