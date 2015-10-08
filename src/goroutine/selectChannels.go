package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		// c 发送数据
		case c <- x:
			x, y = y, x+y
		// 接收quit数据
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)

		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
