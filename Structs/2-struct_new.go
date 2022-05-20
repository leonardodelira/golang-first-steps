package main

import "fmt"

type Pessoa struct {
	Name string
	Idade int32
}

func main(){
	pessoa := new(Pessoa) //Ponteiro
	var pessoa2 Pessoa

	fmt.Println(pessoa)
	fmt.Println(pessoa2)
}