package main

import "fmt"

func main() {
    var ui8 uint8
    var i8 int8
    var i int

    // ui8 = -10 // isto gerará um erro, pois uint8 só aceita valores entre 0 e 255
    // i8 = 128  // isto também gerará um erro, pois int8 só aceita valores entre -128 e 127
    i = 31415

		fmt.Println(ui8, i8, i)
}

/**
uint8       conjunto dos inteiros sem sinal de 8-bits  (0 to 255)
uint16      conjunto dos inteiros sem sinal de 16-bits (0 to 65535)
uint32      conjunto dos inteiros sem sinal de 32-bits (0 to 4294967295)
uint64      conjunto dos inteiros sem sinal de 64-bits (0 to 18446744073709551615)

int8        conjunto dos inteiros com sinal de 8-bits  (-128 to 127)
int16       conjunto dos inteiros com sinal de 16-bits (-32768 to 32767)
int32       conjunto dos inteiros com sinal de 32-bits (-2147483648 to 2147483647)
int64       conjunto dos inteiros com sinal de 64-bits (-9223372036854775808 to 9223372036854775807)

float32     conjunto dos números com ponto-flutuante IEEE-754 de 32-bits
float64     conjunto dos números com ponto-flutuante IEEE-754 de 64-bits

complex64   conjunto dos números com plexos com parte real float32 e parte imaginária
complex128  conjunto dos números com plexos com parte real float64 e parte imaginária

byte        apelido para uint8
rune        apelido para int32
**/