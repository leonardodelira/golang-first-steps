package main

import "fmt"

type S string

var String S = "@LiraLeo"

func main() {
	var text string

	text = "My text!"
	mypicture := "@PhotographLira"
	
	fmt.Println(String)
	fmt.Println(text)
	fmt.Println(mypicture)
}