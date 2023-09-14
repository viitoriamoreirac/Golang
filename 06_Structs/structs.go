package main

import "fmt"

type usuario struct {
	nome		string
	idade		uint8
	endereco	endereco
}

type endereco struct{
	logradouro 	string
	numero		uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var usuario1 usuario
	usuario1.nome = "Vitória"
	usuario1.idade = 20

	fmt.Println(usuario1)

	enderecoExemplo := endereco{"Rua miguel oliveira", 70}
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	// passando apenas um valor
	usuario3 := usuario{nome: "Gustavo"}
	fmt.Println(usuario3)
}