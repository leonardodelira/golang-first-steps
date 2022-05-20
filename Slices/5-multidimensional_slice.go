package main

import "fmt"

func main(){
	a := [][]int {
		{10, 20, 30},
		{1, 2},
		{90, 91},
	}

	for index, value := range a {
		fmt.Println(index, value)
		for _, b := range value {
			fmt.Println(b)
		}
	}
}