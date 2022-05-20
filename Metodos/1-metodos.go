package main

import (
	"fmt"
)

type Pessoa struct {
    Nome string
    Idade int
}

func (*Pessoa) Falar(frases ...interface{}) {
    fmt.Println(frases)
}

func (p *Pessoa) FicouMaisVelho() {
    p.Idade++
}

func main() {
    joaquim := Pessoa{"Carlos", 23}
    joaquim.Falar("Minha idade é:", joaquim.Idade)

    joaquim.FicouMaisVelho()
    joaquim.Falar("Minha idade é:", joaquim.Idade)
}