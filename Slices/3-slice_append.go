package main

import "fmt"

func main(){
	a := make([]int, 5)
	fmt.Println(a)
	a = append(a, 10, 20, 30) //sempre adiciona elementos no final do slice
	fmt.Println(a)
}