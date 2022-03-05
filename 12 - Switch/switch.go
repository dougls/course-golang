package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número de dia inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return "Errado"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(8)
	fmt.Println(dia)

	dia2 := diaDaSemana2(2)
	fmt.Println(dia2)
}
