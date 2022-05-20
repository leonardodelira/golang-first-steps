package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("Divisão por 0!")
	}
	return a / b, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()


	answer, err := divide(10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}    