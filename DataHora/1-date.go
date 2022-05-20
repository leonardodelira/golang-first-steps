package main

import (
	"fmt"
	"time"
)

func main(){
	tm := time.Now()
	fmt.Printf("%s", tm.Local())
}