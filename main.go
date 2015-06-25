package main

import (
	"fmt"

	"./models"
)

/**
 * add function
 */
func add(x int, y int) int {
	return x + y
}

/**
 * first hello world
 */
func main() {
	//	fmt.Println()
	sss := models.GetName()
	fmt.Println(sss)
	fmt.Println(add(2, 3))
	fmt.Println("Hello World!")
}
