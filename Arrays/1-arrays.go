package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3} //array com tamanho fixo
	b := []string{"Abacate", "Sedex", "Leite"} //sem tamanho fixo
	fmt.Println(a, b)

	c := []int{200,300,400}
	fmt.Println(c[1])
	c[1] = 50
	fmt.Println(c[1])
}