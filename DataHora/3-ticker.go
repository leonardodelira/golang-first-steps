package main

import (
	"fmt"
	"time"
) 

func main() { 
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Ticking....")
	}
}

/*
Frequentemente, há casos em que desejamos realizar repetidamente uma determinada tarefa após um determinado intervalo de tempo. Em Golang, conseguimos isso com a ajuda de tickers.
Podemos usá-los com goroutines também para que possamos executar essas tarefas no segundo plano de nosso aplicativo sem interromper o fluxo do aplicativo.
A função que usamos em tickers é a função NewTicker() que leva tempo como um argumento e podemos fornecer segundos e até milissegundos nela.
*/