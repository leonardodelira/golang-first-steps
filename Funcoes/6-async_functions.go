package main

import (
	"fmt"
	"time"
)

func main(){
	go say("World", 1)
	say("Hello", 2)
}

func say(msg string, v time.Duration) {
	time.Sleep(v * time.Second)
	fmt.Println(msg)
}
/*
Uma função pode ser executada de forma assíncrona quando é precedida pelo statement go, que indica que uma função é executada como uma goroutine.
A função main() é executada como uma goroutine por padrão, portanto, todo programa possui ao menos uma goroutine rodando.
*/