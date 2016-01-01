package main

import (
	"fmt"
	"time"
)

func main() {

	aa := time.Date(2015, 11, 3, 18, 47, 40, 0, time.UTC)
	fmt.Println(aa.Format("2006-01-2 15:04:05"))

	now := time.Now()
	fmt.Println(now.Hour())
}
