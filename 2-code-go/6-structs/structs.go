package main

import (
	"fmt"
)

type user struct {
	name       string
	age        uint8
	adressUser adress
}

type adress struct {
	street string
	number uint8
}

func main() {

	var u user
	fmt.Println(u)
	u.name = "Lucas"
	u.age = 22
	fmt.Printf("User full  %+v\n", u)

	u2 := user{"Davi", 23, adress{"Rua 1", 123}}
	fmt.Println(u2)

	exempleadress := adress{"Rua 2", 2}

	u9 := user{"matheus", 35, exempleadress}
	fmt.Println(u9)

	u3 := user{name: "Caitano"}
	fmt.Printf("User 3 full  %+v\n", u3)

	u4 := user{age: 24}
	fmt.Println(u4)

}
