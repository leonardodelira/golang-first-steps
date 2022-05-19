package main

import "fmt"

func main(){
	a := make([]int, 5) //inicia o slice de tamanho 5 preenchendo com 0 os indices
	fmt.Println(a)
	a = append(a, 1, 2, 3, 4)
	fmt.Println(a)
	a[0] = 99
	a[1] = 98
	a[2] = 97
	fmt.Println(a)
}