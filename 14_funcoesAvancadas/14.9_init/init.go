package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}

//função init pode ser repetida no mesmo pacote, uma por arquivo
// é inicializada antes da função main