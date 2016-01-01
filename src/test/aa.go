package main

import (
	"comm/conf"
	"fmt"
)

type Test struct {
	Age  int
	Cats []string
	Pi   float64
}

func main() {
	str := `
  age = 25
  cat = ["12","12"]
  pi = 1.223
  `
	var a = Test{}
	conf.DecodeToml(str, &a)
	fmt.Println(a)
}
