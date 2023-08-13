package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int  // tamanho pré determinado, desvantagem
	fmt.Println(array1)

	var slice []int
	fmt.Println(slice)

	slice = append(slice, 1)
	fmt.Println(slice)

	// arrays internos
	fmt.Println("---------------")

	slice3 := make ([] float32, 10, 12) // cria um slice com capacidade 12 mas salva um array com 10 itens na variável
	fmt.Println(slice3)

	
}