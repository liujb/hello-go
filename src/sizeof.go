package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func getVarAndSize(value interface{}) {
	typeName := reflect.TypeOf(value)
	fmt.Println("type:", typeName, ", value:", value, ", size: ", unsafe.Sizeof(value))
}

func main() {

	var unum8 uint8 = 1
	var unum16 uint16 = 1
	var unum32 uint32 = 1
	var unum64 uint64 = 1
	var num8 int8 = 1
	var num int = 1
	var num32 int32 = 1
	var num64 int64 = 1
	var float1 float32 = 1.32
	var float2 float64 = 1.32
	var b bool = true
	var c rune = 'a'
	var str string = "aaaa"
	var com complex64 = 1.2
	var com2 complex128 = 1.2

	slice := []string{"sd", "sd"}

	mapT := make(map[string]interface{})
	mapT["a"] = unum8
	mapT["b"] = unum16
	mapT["c"] = unum32
	mapT["d"] = unum64
	mapT["e"] = num8
	mapT["f"] = num32
	mapT["f"] = num64
	mapT["h"] = float1
	mapT["i"] = float2
	mapT["ss"] = num

	mapT["j"] = c
	mapT["k"] = str
	mapT["l"] = b
	mapT["m"] = com
	mapT["n"] = com2
	mapT["o"] = slice

	// for _, v := range mapT {
	// 	getVarAndSize(v)
	// }
	fmt.Println(unsafe.Sizeof(c))

}
