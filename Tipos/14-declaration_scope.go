package main

import "fmt"

func main() {
    var a bool
    for {
        if a {
            break
        } else {
            b := 1
            a = true
        }
    }

    fmt.Println(b) // Variável b não foi declarada no escopo atual nem nos escopos superiores
}