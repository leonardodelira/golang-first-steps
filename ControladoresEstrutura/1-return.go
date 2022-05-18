package main

import "fmt"

func Lambda(name string) string {
	return fmt.Sprint("Hi ", name)
}

func main(){
	println(Lambda("Lira"))
	return

	//fmt.Println("teste") Esse bloco de código não irá ser executado por conta do return na linha 11
}