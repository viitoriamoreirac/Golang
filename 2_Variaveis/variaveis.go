package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 int = 45
	)

	fmt.Println(variavel3, variavel4)
}