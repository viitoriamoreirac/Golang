package main

import (
	"errors"
	"fmt"
)

func main() {
	var tipoInteiro int = 5
	var tipofloat float32 = 5.23
	var tipoFloat float64 = 784.23
	var tipoString string = "tipos"

	// apenas alguns exemplos, não copiei a aula

	fmt.Println(tipoInteiro)
	fmt.Println(tipofloat)
	fmt.Println(tipoFloat)
	fmt.Println(tipoString)


	var stringVazia string 
	var intVazio int 

	fmt.Println(stringVazia) // imprime uma string vazia
	fmt.Println(intVazio) // imprime o número zero

	var booleano1 bool  // sem valor atribuido, logo, falso
	var booleano2 bool = true
	
	fmt.Println(booleano1)
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro) // imprime <nil>, valor vazio para erro

	var erro2 error = errors.New("Novo erro") // errors é o pacote de erro em go
	fmt.Println(erro2)
}