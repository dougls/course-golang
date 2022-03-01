package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 100
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Fim numeros inteiros

	// números reais

	var numeroReal1 float32 = 123.64
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.58
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Fim numeros reais

	// Strings
	var str string = "lolololo"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Fim Strings

	texto := 5
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
