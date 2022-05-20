package main

import (
	"fmt"
)
  
func main() {

    a := 10
    b := "pablo"

    var c *int
    c = &a //com "&" pegamos o endereco de memória da variavel "a", não o seu valor que seria "10"

    d := &b

    fmt.Println(a, c)
    fmt.Println(b, d)
}

/*
	Ponteiro é uma variavel para armazenar o endereco de memória de outra variavel
		Além disso, o ponteiro aponta a onde a memória está localizada e nos permite acessar o dado armazenado naquele endereco de memória
	*/