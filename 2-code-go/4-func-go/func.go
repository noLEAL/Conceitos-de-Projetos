package main

import (
	"fmt"
)

// () - Serve para passar parâmetros para a função, sempre colocando o tipo do parâmetro
// depois do parenteses, serve para espesificar o tipo de dados que a função ira retornar
// {} - Serve para colocar o corpo da função
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// () - Serve para passar parâmetros para a função, sempre colocando o tipo do parâmetro
// proximo () - Serve para espesificar o tipo de dados que a função ira retornar,
//
//	sinalizando que irar retornar 4 resultados diferentes pois exitem 4 retornos diferentes
func calcMath(num1, num2 int8) (int8, int8) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}

func main() {

	sum := somar(10, 20)
	fmt.Println(sum)

	// declarando uma variavel como uma função
	var funcVar = func(str string) string {
		var str2 string = " - Dentro da função"
		fmt.Println(str, str2)
		return str
	}

	result := funcVar("Passando str como parametro")
	fmt.Println(result)

	sum, sub := calcMath(10, 15)
	fmt.Println(sum, sub)

	// Ignorando um retorno
	// Se eu quiser ignorar um retorno, eu posso usar o _ (underline)
	ret1, ret2 := calcMath(10, 15)
	fmt.Println(ret1, ret2)

}
