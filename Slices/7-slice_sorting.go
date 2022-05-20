package main

import (
	"fmt"
	"sort"
)

func main(){
		Autor := []struct{
			name string
			article int	
			id int
		}{
			{"Mina", 303, 1},
			{"Cina", 403, 2},
			{"Tina", 503, 3},
			{"Bina", 603, 4},
		}

		sort.Slice(Autor, func(p, q int) bool {
			return Autor[p].name < Autor[q].name
		})

		fmt.Println("Sort Autor de acordo com o campo nome:")
		fmt.Println(Autor)

		valores := []int{10,22,3,43,5,62,7}
		sort.Ints(valores)
		fmt.Println(valores)

		sort.Sort(sort.Reverse(sort.IntSlice(valores)))
		fmt.Println(valores)
}