package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e Slice")

	var array1 [5]string
	array1[0] = "Posição"
	fmt.Println(array1)

	array2 := [5]string{"Posição 2", "outra strign"}
	fmt.Println(array2)

	// [...] fixa o tamanho do array conforme a quantidade de item que vc determinou apos o tipo
	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	//Diferença entre array e slice é o fato de declarar um tamanho ou não para esse objeto

	//slice = tamanho dinamico, aumenta conforme apend

	slice := []int{1, 5, 6, 8, 5}

	slice = append(slice, 1000)

	fmt.Println(slice)

	//slice Não é um array, mas usa um como estrutura  -  slice=fatia em ingles

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	//======

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = 666666
	fmt.Println(slice2)

}
