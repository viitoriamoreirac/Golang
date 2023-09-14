package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1  // COPIA DO VALOR ≠ PONTEIRO

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO => REFERENCIA DE MEMORIA - ENDEREÇO DE MEMORIA ONDE A VARIAVEL ESTÁ SALVA
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3  // valor de memoria

	fmt.Println(variavel3, ponteiro, *ponteiro) // com asterisco = desreferenciação
}