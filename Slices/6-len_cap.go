package main

import "fmt"

func main(){
	a := make([]int, 5, 10)
	fmt.Println(a)
	fmt.Println("len ", len(a))
	fmt.Println("cap ", cap(a))

	a = append(a, 10, 20, 30, 40, 50)
	fmt.Println(a)
	fmt.Println("len ", len(a))
	fmt.Println("cap ", cap(a))

	a = append(a, 60)
	fmt.Println(a)
	fmt.Println("len ", len(a))
	fmt.Println("cap ", cap(a))
}

/*
Um slice é uma abstração que usa um array nos bastidores.
A função cap() informa a capacidade do array subjacente. len() informa quantos itens estão no array.
*/