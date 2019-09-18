package main

import "fmt"

func soma(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func main() {
	numero1 := 10
	numero2 := 20

	resultado := soma(numero1, numero2)
	resultado *= 2
	fmt.Println(resultado)
}
