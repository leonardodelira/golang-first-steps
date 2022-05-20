package main

import (
	"fmt"
	"time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Local())

    oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
    diff := now.Sub(oldTime)
    //diff é do tipo Duration
    
    //Em horas
    fmt.Printf("Hours: %f\n", diff.Hours())

    //Em minutos
    fmt.Printf("Minutes: %f\n", diff.Minutes())

    //Em segundos
    fmt.Printf("Seconds: %f\n", diff.Seconds())
}