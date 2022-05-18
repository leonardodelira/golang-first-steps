package main

import "fmt"

type MyStruct struct {
  Name string
}

var (
  ch1 = make(chan *MyStruct, 1)
	ch2 = make(chan int, 2)
	ch3 = make(chan string, 3)
)

func main() {
	ch1 <- &MyStruct{Name: "Lira"}
	defer close(ch1)
	fmt.Println(ch1, &ch1, <-ch1)

	ch2 <- 10
	ch2 <- 11
	//ch2 <- 11 Neste caso daria erro pq eu só liberei 2 canais para este channel
	defer close(ch2)
	fmt.Println(<-ch2) //Os valores são exibidos em ordem na qual foram atribuidos ao channel
	fmt.Println(<-ch2)

	ch3 <- "One"
	ch3 <- "Two"
	ch3 <- "Three"
	defer close(ch3)
	fmt.Println(<-ch3) //Os valores são exibidos em ordem na qual foram atribuidos ao channel
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
}