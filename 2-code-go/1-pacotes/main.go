package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Meu primeiro modulo")
	auxiliar.Escrever()

	
	error := checkmail.ValidateFormat("brunopazet@gmail.com")
	fmt.Println(error)
}
