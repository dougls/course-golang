package main

import "fmt"

//maneira 1 de como declarar variáveis
func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//maneira 2 de como declarar variáveis
	var (
		variavel3 string = "lalalala"
		variavel4 string = "lelele"
	)

	fmt.Println(variavel3, variavel4)

	//outra maneira de declarar variáveis
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	//constante tem o mesmo proposito de variáveis
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//trocando o valor das variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
