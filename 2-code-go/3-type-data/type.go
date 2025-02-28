package main

import (
	"errors"
	"fmt"
)

func main() {

	//int8, int16, int32, int64
	//int = vai usar arquitetura do computador como base
	var num int16 = 1000
	fmt.Println(num)

	//uint8, uint16, uint32, uint64 - unsygned int - não aceita negativos
	var num2 uint32 = 100
	fmt.Println(num2)

	//alias
	// INT32 = RUNE
	var num3 rune = 123
	fmt.Println(num3)

	// BYTE = UINT8
	var num4 byte = 123
	fmt.Println(num4)

	//FLOAT32, FLOAT64
	//FLOAT = Usa arquitetura do computador como base para tipar a variável
	var num5 float32 = 123.45
	fmt.Println(num5)

	//STRING, sempre usar aspas duplas quando for string
	var str string = "texto"
	fmt.Println(str)

	//CHAR = NÃO EXISTE EM GO, MAS PODEMOS USAR O BYTE QUE REPRESENTA O CHAR NA TABELA ASCII
	//IRA IMPRIMIR O VALOR DECIMAL 66 QUE CORRESPONDE A LETRA B NA TABELA ASCII
	char := 'B'
	fmt.Println(char)

	//Valor zero de uma variável é o valor que a variável recebe quando não é atribuido nenhum valor
	//fazendo com que ela imprima uma string vazia ou um numero 0
	var str2 string
	var num6 int
	fmt.Println(str2, num6)

	//BOOLEAN
	//true ou false
	//valor zero é false
	var boolean bool = true
	fmt.Println(boolean)

	//ERRO
	//tipo de dado que representa um erro
	//valor zero é nil
	var Varerro error = errors.New("Erro interno")
	fmt.Println(Varerro)
	//não é um string, e sim uma variavel de tipo error
}
