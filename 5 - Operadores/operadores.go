package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Fim dos aritméticos

	// Atribuição
	var variavel1 string = "String" // modo 1 declaração variável
	variavel2 := "String2"          // modo 2 declaração variável
	fmt.Println(variavel1, variavel2)

	// Fim das atribuições

	// Operadores Relacionais(sempre retornam boolean)
	fmt.Println(1 > 2)  // maior
	fmt.Println(1 >= 2) // maior ou igual
	fmt.Println(1 == 2) // igual à
	fmt.Println(1 <= 2) // menor ou igual
	fmt.Println(1 < 2)  //  menor
	fmt.Println(1 != 2) // diferente
	// Fim dos operadores relacionais

	// Operadores lógicos
	fmt.Println("------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // and
	fmt.Println(verdadeiro || falso) // or
	fmt.Println(!verdadeiro)         // negação
	fmt.Println(!falso)              // negação
	// Fim dos operadores lógicos

	// Operadores unários
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	// Fim dos operadores unários

	// Operadores Tenários

}
