package main

import "fmt"

func main() {
	defer fmt.Println("World") //com a anotacao defer, essa linha de código será executada no fim da funcao
	fmt.Println("Hello")
}