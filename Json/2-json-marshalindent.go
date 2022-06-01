package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ApiLogin struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func main() {

	a := ApiLogin{"Jefferson", "033.343.434-89"}
	// improving output for json format viewing
	json, err := json.MarshalIndent(a, "", "\t") //Formata a saída do JSON
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}