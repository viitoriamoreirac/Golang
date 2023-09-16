package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fezAniversario() {
	u.idade++
}

func main () {
	usuario1 := usuario{"vitoria", 20}
	usuario1.salvar() // metodo
	fmt.Println(usuario1.maiorDeIdade()) // metodo q testa maioridade

	usuario2 := usuario{"eliane", 47}
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade()) // metodo q testa maioridade
	usuario2.fezAniversario() // metodo p aumentar idade
	fmt.Println(usuario2.idade) // imprimindo idade nova
}