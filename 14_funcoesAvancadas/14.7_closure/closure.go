package main

import "fmt"

func closure() func(){
	text := "dentro de closure"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Dentro de main"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}