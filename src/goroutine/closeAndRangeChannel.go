package main

import (
	"fmt"
	"runtime"
)

/**
 * fibonacci数列
 * @param  {[type]} n int [description]
 * @param  {[type]} c chan int [description]
 * @return {[type]}   [description]
 */
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	fmt.Println(runtime.NumCPU())

	c := make(chan int, 20)
	go fibonacci(cap(c), c)

	// for i := range c能不断读取channel里面的数据
	// 直到该channel关闭
	if v, ok := <-c; ok {
		// 测试channel是否被关闭
		fmt.Println(ok, v) // true 1
		for i := range c {
			fmt.Println(i)
		}
	} else {
		fmt.Println("aaa")
	}

}
