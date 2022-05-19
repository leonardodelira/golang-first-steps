package main

import "fmt"

/*
Em GO, a função também é um tipo. Duas funções são do mesmo tipo se tiverem os mesmos argumentos e os mesmos valores de retorno. Ao passar uma função como um argumento para outra função, a assinatura exata da função deve ser especificada na lista de argumentos.
*/
func main() {
	fmt.Println("somar", executarOperacao(somar,1,2,3))
	fmt.Println("multiplicar", executarOperacao(multiplicar,1,2,5))
}

func somar(a, b int, c ...int) int {
	total := a + b
	for _, value := range c {
		total += value
	}
	return total
}

func multiplicar(a, b int, c ...int) int {
	total := a * b;
	for _, value := range c {
		total *= value
	}
	return total
}

func executarOperacao(f func(int, int, ...int) int, a int, b int,c int) int {
	return f(a, b, c)
}