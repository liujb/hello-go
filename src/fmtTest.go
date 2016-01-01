package main

import (
	"fmt"
	"strconv"
)

func main() {
	abc := "abc"
	fmt.Sprintf("aaa", abc)
	fmt.Println(abc)

	aa := 12
	fmt.Printf("%x\n", aa)

	b := 0x80
	fmt.Printf("%d\n", b)

	str := "0x0c"
	i, _ := strconv.ParseInt(str, 0, 64)
	i2, _ := strconv.ParseInt("80", 16, 64)
	fmt.Println(i)
	fmt.Println(i2)

	fmt.Println("\r\n")

	f := 12.33
	// strFl := fmt.Sprintf("%f", f)
	strFl := fmt.Sprint(f)
	fmt.Println(strFl)

	str2 := "001.12"
	ff, _ := strconv.ParseFloat(str2, 64)
	fmt.Print(ff)
}

// 0 1 2 3 4 5 6 7 8 9 a b c d e f
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
