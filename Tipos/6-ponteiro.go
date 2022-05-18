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
	u := &trabalhador //Ao passar o &, pego a referencia em memória da variavel em questão

	trabalhador.Idade = 20
	fmt.Println("u: ", u)

	u.Pessoa.Idade = 21
	fmt.Println("trabalhador: ", trabalhador)
}