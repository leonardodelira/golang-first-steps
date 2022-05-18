package main

import "fmt"

type Carro struct {
	Modelo string
	Valor int32
}

func main(){
	var m1 map[string]int
	var m2 = make(map[string]int)
	var m3 = map[string]string{"name": "Lira"}
	
	var m4 = m3
	m4["age"] = "Vinte"

	var m5 map[string]string
	m5 = make(map[string]string)

	var m6 = map[string]Carro{"primeiro": {"Astra", 21000}}

	fmt.Println(m1, m2, m3, m4, m5, m6)
	fmt.Println(len(m1), len(m2), len(m3), len(m4), len(m5), len(m6))
}