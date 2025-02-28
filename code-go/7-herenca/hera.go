package main

import "fmt"

type person struct {
	name     string
	age      uint8
	height   uint8
	lastname string
}

type student struct {
	person     //Para trabalhar com "herença" quando referenciar outro objeto, é precisso não colocar o tipo desse objeto
	couses     string
	university string
}

func main() {

	fmt.Println("Herança")

	p1 := person{"João", 20, 178, "Pedro"}
	fmt.Println(p1)

	e1 := student{p1, "Engenharia", "UPS"}
	fmt.Println(e1)

	fmt.Println(e1.height)
	fmt.Println(e1.name)
	fmt.Println(e1.couses)

}
