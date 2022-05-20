package main

import (
	"fmt"
	"time"
) 

func main() { 
	fmt.Println("Iniciou ", time.Now().Local())
	time.Sleep(2 * time.Second)
	fmt.Println("Após 2 segundos", time.Now().Local())
}