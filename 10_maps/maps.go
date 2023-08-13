package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string] string {
		"nome": 		"pedro",
		"sobrenome": 	"silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string] string{
		"nome": {
			"primeiro": 	"João",
			"ultimo": 		"Lopes",
		},
	}

	fmt.Println(usuario2)
}