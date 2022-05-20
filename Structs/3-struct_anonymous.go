package main

import "fmt"

func main(){
	pessoa := struct{
		Name string
		Idade int32
	}{
		Name: "Lira",
		Idade: 24,
	}

	fmt.Println(pessoa)
}