package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i <= 5; i++ {
		runtime.Gosched()
		fmt.Println(i, s)
	}
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(2))
	// 通过关键字go启动 goroutine
	go say("world")
	say("hello")
}
