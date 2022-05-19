package main

import (
	"fmt"
	"time"
)

func main(){
	contagemRegressiva(10)
}

func contagemRegressiva(v int) {
	fmt.Println("Tempo: ", v)
	
	if v == 0 {
		return
	}	
	
	time.Sleep(1 * time.Second)
	contagemRegressiva(v - 1)
}