package main

import "fmt"

//interface{} signifca que qualquer valor pode ser atribuido a variavel

type MyStruct struct {
	Msg interface{}
}

func main(){
	structure := MyStruct{}
	structure.Msg = "Lira"
	fmt.Println(structure.Msg)

	structure.Msg = map[int]string{1: "um", 2: "dois", 3: "tres"}
	fmt.Println(structure.Msg)
}