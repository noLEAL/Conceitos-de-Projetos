package main

import (
	"fmt"
)

func main() {
	soma := 1 + 2
	subtração := 1 - 2
	div := 1 / 2
	mutiplicacao := 1 * 2
	restoDaDivisao := 1 % 2

	fmt.Println(soma, subtração, div, mutiplicacao, restoDaDivisao)

	//GO é fotemente tipado, então não é possivel fazer operações com tipos diferentes ex: int8 + int16
	// var num1 int8 = 10    NÃO PODE
	// var num2 int16 = 20           NÃO PODE
	// soma := num1 + num2                   NÃO PODE
	// fmt.Println(soma)                             NÃO PODE

	//ATRIBUIÇÃO

	var variavel1 string = "String" //Precisa declarar o tipo da variavel
	variavel2 := "String2"          //Não precisa declarar o tipo da variavel, o GO infere o tipo da variavel
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS - Sempre retornam um valor booleano (True ou False)
	fmt.Println(1 > 2)  // > - Maior
	fmt.Println(1 < 2)  // < - Menor
	fmt.Println(1 >= 2) // >= - Maior ou igual
	fmt.Println(1 <= 2) // <= - Menor ou igual
	fmt.Println(1 == 2) // == - Igual
	fmt.Println(1 != 2) // != - Diferente

	//OPERADORES LÓGICOS
	fmt.Println(true && false) // && - E
	fmt.Println(true || false) // || - OU
	fmt.Println(!true)         // ! - Negação

	//OPERADORES UNÁRIOS
	num := 10 //Atribuição
	num++     //Incremento           num = num + 1
	num += 15 //Soma                 num = num + 15
	num--     //Decremento           num = num - 1
	num -= 20 //Subtração            num = num - 20
	num *= 3  //Multiplicação        num = num * 3
	num /= 2  //Divisão              num = num / 2
	num %= 2  //Resto da divisão     num = num % 2
	fmt.Println(num)

	//OPERADORES TERNÁRIOS
	//Não existe em GO

	//texto := numero > 5 ? "Maior que 5" : "Menor que 5" //Não existe em GO alternativa abaixo

	numero := 10
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
