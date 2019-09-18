package main

import "fmt"

func main() {
	var code1 bool
	code1 = false

	if code1 {
		fmt.Println("Passou")
	} else {
		fmt.Println("Não passou")
	}

	//OU

	if code2 := true; code2 {
		fmt.Println("Passou no teste")
	} else {
		fmt.Println("Não passou no teste")
	}

}
