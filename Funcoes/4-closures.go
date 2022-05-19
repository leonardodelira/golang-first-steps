package main

import "fmt"

func main(){
	contador := contar()

	fmt.Println(contador())
	fmt.Println(contador())
	fmt.Println(contador())
	fmt.Println(contador())
	fmt.Println(contador())
}

func contar() func() int {
	var i int

	return func() int {
		i++
		return i
	}
}