package main

import "fmt"

type Pessoa struct {
	Name string
	Idade int32
}

func main(){
	pessoa := Pessoa{Name: "Lira", Idade: 24}
	fmt.Println(pessoa)
}