package main

import "fmt"

type MyStruct struct {
	nick string
}

func main(){
	langs := []string{"Erlang", "Elixir", "Javascript", "Java"}
	fmt.Println(langs)
	for key, value := range langs {
		fmt.Println(key, value)
	}

	countryCapital := map[string]string{"Brasil": "Brasilia", "EUA": "Washington", "France": "Paris"}
	for country, capital := range countryCapital {
		fmt.Println(country, capital)
	}

	jobs := make(chan int, 3)
  for j := 1; j <= 3; j++ {
    jobs <- j
  }
  close(jobs)

	a := []MyStruct{{"@devopsbr"}, {"@go_br"}, {"@awsbrasil"}, {"@go_br"}, {"@devopsbh"}}
  for i, v := range a {
    fmt.Println(i, v.nick)
  }

	var key string
  var val interface{}
  m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
  for key, val = range m {
    fmt.Println(key, val)
  }
}