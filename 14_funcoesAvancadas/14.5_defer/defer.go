package main

import "fmt"

func func1() {
	fmt.Println("Executando função 1")
}

func func2() {
	fmt.Println("Executando a função 2")
}

func main() {
	defer func1() // defer = adiar - adia a execução
	func2()
}