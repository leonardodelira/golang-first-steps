package main

import "fmt"

func main() {
		for n := 1; n < 10; n++ {
			if n % 2 == 0 {
				goto x
			}
		}
    x:
			y := "Hello"
    	fmt.Println(y)
}