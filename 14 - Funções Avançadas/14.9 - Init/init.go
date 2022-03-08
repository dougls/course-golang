package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a função init") // será executada sempre em primeiro lugar, independente de sua posição
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada") // função main é o start do programa
	fmt.Println(n)
}
