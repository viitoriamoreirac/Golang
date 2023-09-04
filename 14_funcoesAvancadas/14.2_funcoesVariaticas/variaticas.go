package main

import "fmt"

func soma(numeros ...int) int{
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalSoma := soma(1, 2, 4, 8, 16)
	fmt.Println(totalSoma)
}