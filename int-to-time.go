package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	t := time.Unix(1435325220, 0)

	fmt.Println(t.Format("2006_01_02_15"))
	fmt.Println(t)
}
