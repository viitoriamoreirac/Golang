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


}