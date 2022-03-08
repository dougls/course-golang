package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do %s com idade de %d anos no banco de dados\n", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Douglas", 29}
	usuario1.salvar()

	usuario2 := usuario{"Leticia", 25}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

}
