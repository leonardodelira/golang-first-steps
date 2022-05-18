package main

func main() {
  n := 5
  for n > 0 {
    n--
    println(n)
  }

  // declara i e acrescenta i em 1 a cada iteração
  for i := 0; i < 5; i++ {
    println(i)
  }

  n = 5
  for i := 0; i < n; i++ {
    if i <= 2 {
      continue
    } else {
      println("i > 2 = ", i)
    }
  }

  n = 5
  for i := 0; i < n; i++ {
    if i == 2 {
      break
    } else {
      println("i: ", i)
    }
  }
}	