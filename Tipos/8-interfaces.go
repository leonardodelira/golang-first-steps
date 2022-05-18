package main

import "fmt"

func Printar(v ...interface{}) {
    fmt.Println(v...)
}

var Msg interface{}

func main() {
    Printar(18.4, "Golang")
}