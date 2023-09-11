package main

import "fmt"

func func1() {
	fmt.Println("Executando função 1")
}

func func2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} 
	return false
}
func main() {
	defer func1() // defer = adiar - adia a execução
	func2()
	fmt.Println(alunoEstaAprovado(7, 8))
}

/* pode ser muito útil lidando com banco de dados, consulta, inserção no banco
insert e update, fecha conexão com banco de dados usando defer, independente do return da função. */