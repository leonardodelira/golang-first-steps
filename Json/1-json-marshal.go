package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Name struct {
	Name string `json:"name"`
	Document string `json:"document"`
}

func main() {
	person := Name{"Leonardo", "458.000.123-87"}
	fmt.Println(person)

	m, err := json.Marshal(person) //transforma a struct em json
	if err != nil {
		log.Println(err)
	}
	fmt.Println(m) //show bytes

	fmt.Println(string(m)) //show string json
}