package main

import (
	"errors"
	"fmt"
)

//error customizado --> New
func divide(x int, y int) (int, error) {
    if y == 0 || x == 0{
        return 0, errors.New("Divisão por 0!")
    }
    return x/y, nil
}

func main() {
    answer, err := divide(5,0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(answer)
}