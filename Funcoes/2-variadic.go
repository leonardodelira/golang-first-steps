package main

import "fmt"

func main() {
	total := multiplesParams(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)
}

func multiplesParams(a int, b ...int) int {
	var total int
	for _, number := range b {
		total += number
	}
	return total + a
}