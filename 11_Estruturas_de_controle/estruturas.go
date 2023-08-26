package main

import "fmt"

func main() {
	fmt.Println("Estruturas")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Maior,0")
	} else {
		fmt.Println("Menor,0")
	}
}