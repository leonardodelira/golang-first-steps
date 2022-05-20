package main

import (
	"fmt"
)

type Falador interface {
    Falar(...interface{})
		Calado()
}

type Pessoa struct {
    Nome string
    Idade int
}

//De forma implicita, a struct Pessoa está implementando a interface Falador
func (Pessoa) Falar(frases ...interface{}) {
    fmt.Println(frases...)
}

func (Pessoa) Calado() {
	fmt.Println("...")
}

func (p *Pessoa) FicarMaisVelho() {
    p.Idade++
}

func falar(falador Falador, frases ...interface{}) {
    falador.Falar(frases...)
}

func main() {
	fulano := &Pessoa{"Carlos", 23}
	fulano.Falar("Minha idade é:", fulano.Idade)

	fulano.FicarMaisVelho()
	falar(*fulano, "Minha idade é:", fulano.Idade) // Nenhum erro pois Pessoa implementa a interface Falador

	fulano.Calado()
}