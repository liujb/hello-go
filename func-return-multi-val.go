package main

import "fmt"


func multi() (string, int){
    return "nihao", 10
}
func main(){
    x, y := multi();
    fmt.Println(x);
    fmt.Println(y);
    fmt.Println("gogoog");
}
