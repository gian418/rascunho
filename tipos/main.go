package main

import "fmt"

// Pessoa é um tipo que herda do tipo strings
type Pessoa string

func main() {
	var nome string
	var idade int
	var dinheiro float32
	var verificar bool
	var albano Pessoa

	nome = "Gian Haack"
	idade = 29
	dinheiro = 56.67
	verificar = true
	albano = "Albano"

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(dinheiro)
	fmt.Println(verificar)
	fmt.Println(albano)
}
