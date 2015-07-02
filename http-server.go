package main

import (
    "net/http"
)

func main(){
    http.ListenAndServe(":8021", nil)        
}
