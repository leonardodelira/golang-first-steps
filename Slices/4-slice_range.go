package main

import "fmt"

func main(){
	a := []int{10, 20, 30, 40}
	for index, value := range a {
		fmt.Println(index, value)
	}
}