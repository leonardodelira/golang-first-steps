package main

import (
	"fmt"
	"time"
)

func main(){
	current_time := time.Now()
	
	fmt.Println("Dia ", current_time.Day())
	fmt.Println("Mês ", current_time.Month())
	fmt.Println("Ano ", current_time.Year())

	fmt.Println("Hora ", current_time.Hour())
	fmt.Println("Minutos ", current_time.Minute())	
	fmt.Println("Segundos ", current_time.Second())

	fmt.Println(current_time.Format("2006-01-02 15:04:05"))
	fmt.Println(current_time.Format("2006-January-02"))
	fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))
}