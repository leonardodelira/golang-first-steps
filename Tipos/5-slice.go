package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int32
}

type Trabalhador struct {
	Cargo string
	Pessoa
}

func main(){
	trabalhador := Trabalhador{"Software Enginner", Pessoa {"Leonardo", 24}}
	fmt.Println(trabalhador)
}