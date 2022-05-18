package main

func main() {
	number := 10
	switch number {
	case 9:
		println("here: 9")
	default:
		println("here: default")
	}

	for i := 0; i<10; i++ {
		switch i {
		case 7:
			goto LABELGOTO
		default:
			println("i: ",i)
		}
	}

	LABELGOTO:
		callfunc()
}	
    
func callfunc() {
	println("callfunc called")
}