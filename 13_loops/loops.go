package main

import (
	"fmt"
	"time"
)

func main()  {
/*	i := 0

	for i<10 {
		i++
		fmt.Println("Incrementando 1")
		time.Sleep(time.Second)
	}

	for i := 0; i < 10; i += 2 {
		fmt.Println("Incrementando 1")
		time.Sleep(time.Second)
	}
*/
	nomes := [3] string{"Vitoria", "maria", "julia"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
		time.Sleep(time.Second)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":			"leonardo",
		"sobrenome":	"silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}