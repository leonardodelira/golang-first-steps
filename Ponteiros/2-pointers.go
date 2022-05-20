package main

import "fmt"

type Pessoa struct {
	Name string
	Idade int32
	Endereco string
}

func (p *Pessoa) maiorDeIdade() bool {
	if p.Idade >= 18{
		return true
	} 
	return false
}

func (p *Pessoa) adicionaIdade() {
	p.Idade++
}

func main(){
	fulano := Pessoa{Name: "Alfredo", Idade: 16, Endereco: "Rua Paulista"}

	for {
		deMaior := fulano.maiorDeIdade()
		if deMaior == true {
			fmt.Println(fulano.Name, " ", fulano.Idade, " anos")
			break
		}
		fmt.Println(fulano.Name, " ", fulano.Idade, " anos")
		fulano.adicionaIdade()
	}
}


/*
Quando usar ponteiros?
	Quando precisamos alterar o valor original da instancia

Permitindo que as variáveis tenham valores nil: Essa é uma pegadinha que leva muitos iniciantes em Go a criar bugs. Uma variável que é definida, mas não tem um valor atribuído, não é avaliada como nulo, mas sim com seu valor padrão. Mas e se você precisar que seja nulo? Então, você usa um ponteiro!
Copiar estrutura com valores muito grandes: Embora a memória não seja um problema tão grande hoje, você ainda pode querer usar uma referência se o valor do objeto for ridiculamente grande e não houver realmente necessidade de repassá-lo o tempo todo.
*/