package main

import "fmt"

func fibo(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	} 
	return fibo(posicao-2) + fibo(posicao-1)
}

func main() {
	fmt.Println("recursivas")
	// 1, 1, 2, 3, 5, 8, 13, 21

	posicao := uint(20)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibo(i))
	}
}