package main

import "fmt"

func main() {

	//==================================================
	var typevar1 string = "value1" // declaração de variavel com tipo, var {nome} {tipo} = {valor}
	infvar2 := "value2"            // inferencia de tipo, se sabe o tipo da variavel pelo valor atribuido

	fmt.Println(typevar1)
	fmt.Println(infvar2)

	//==================================================

	var ( // declaração de multiplas variaveis
		multivar1 = "dajdlkas"
		multivar2 = 123
	)

	fmt.Println(multivar1, multivar2)

	//==================================================

	swap1, swap2 := "value1", "value2" // declaração de multiplas variaveis com inferencia de tipo

	fmt.Printf("Variavel 1: %s Variavel 2: %s \n", swap1, swap2)

	//==================================================

	fmt.Print("Trocando Valores\n")
	swap1, swap2 = swap2, swap1 // troca de valores entre variaveis

	fmt.Printf("Variavel 1: %s Variavel 2: %s \n", swap1, swap2)

	//==================================================
}
