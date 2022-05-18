package main

import (
	"fmt"
	"reflect"
)

func main(){
	comum()
	func1(1)
	func2 := func2(1,2,2.4)
	fmt.Println("func2 ", func2)
	func3 := func3(1,2,3.5)
	fmt.Println("func3 ", func3)
	func4("Valores", 1,2,3,4,5,6)
}

func comum() {
	fmt.Println("Chamada funcao comum")
}

func func1(x int) int {
	fmt.Print(x)
	return x
}

func func2(a, _ int, z float32) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(z)
}

func func3(a, b int, z float32) (bool) {
	return reflect.TypeOf(a) == reflect.TypeOf(z)
}

func func4(prefix string, values ...int) {
	fmt.Println(prefix)
	fmt.Println("founded ", len(values), " values")
	for x := range values {
		fmt.Println(values[x])
	}
}

// func(a, b int, z float64, opt ...interface{}) (success bool)
// func(int, int, float64) (float64, *[]int)
// func(n int) func(p *T)