package main

import (
	"encoding/json"
	"fmt"
)

type Login struct {
	Email string `json:"email"`
	Username string `json:"username,omitempty"` //Caso esse atributo não seja preenchido na struct, no JSON ele será ocultado
	Password string `json:",omitempty"` //Quando queremos que o atributo JSON tenha o mesmo valor da struct, basta não colocar nenhum nome
	Address string `json:"-"` //Campo será ignorado no JSON
	Phone string `json:"telphone,"` //Campo irá aparecer no JSON com o valor "-"
}

func main() {
	l := Login{Email: "leonardo.lira@teste.com", Username: "", Password: "****", Address: "Rua da Bahia", Phone: "19987876524"}
	fmt.Println(l)

	m, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(m))
}