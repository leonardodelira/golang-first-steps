package main

import "fmt"

type R struct{
	R string
}

type IRead interface{
	Read() string
}

func (r *R) Read() string {
	return fmt.Sprintf("Only: call Read")
}

func Call(ir IRead) string {
  return fmt.Sprintf("Read: %s", ir.Read())
}

func main(){
	var iread IRead
	r := R{"hello interface"}
	
	//uma forma de usar interface
	iread = &r
	fmt.Println(iread, r)
	fmt.Println(iread.Read())

	//outra forma de usar interface
	r2 := R{"second call to struct R"}
	fmt.Println(Call(&r2), r2.R)
}