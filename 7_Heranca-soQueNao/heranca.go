package main

import "fmt"

type person struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudant struct {
	person    // essa é a herança, puxando todos os "atributos" de person
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	person1 := person{"Joao", "Pedro", 20, 178}
	fmt.Println(person1)

	estudant1 := estudant{person1, "engenharia", "usp"}
	fmt.Println(estudant1)
}