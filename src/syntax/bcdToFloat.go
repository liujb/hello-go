package main

import (
	"fmt"
	"strconv"
)

func BCDToInt(bcd byte) int {
	tmp := (0xff&(bcd>>4))*10 + (0xf & bcd)
	return int(tmp)
}

func main() {

	bytes := []byte{1, 33, 38, 0, 37}
	bytes2 := []byte{49, 17, 134, 72}
	bytes3 := []byte{21, 17, 3, 24, 71, 64}

	lat := ""
	for k, v := range bytes {
		if k == 2 {
			lat += "."
		}
		lat += strconv.Itoa(BCDToInt(v))
	}
	fmt.Println(lat)

	lng := ""
	for k, v := range bytes2 {
		if k == 1 {
			lng += "."
		}
		lng += strconv.Itoa(BCDToInt(v))
	}
	fmt.Println(lng)

	time := "20" + strconv.Itoa(BCDToInt(bytes3[0])) + "-" + strconv.Itoa(BCDToInt(bytes3[1])) + "-" + strconv.Itoa(BCDToInt(bytes3[2])) + " " + strconv.Itoa(BCDToInt(bytes3[3])) + ":" + strconv.Itoa(BCDToInt(bytes3[4])) + ":" + strconv.Itoa(BCDToInt(bytes3[5]))

	fmt.Println(time)
}
