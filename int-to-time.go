package main

import (
    "fmt"
    "time"
)

func main(){

    t := time.Unix(1435325220, 0)
    fmt.Println(t.Format("2006_01_02_15"))
    fmt.Println(t)

}

