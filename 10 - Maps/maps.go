package main

import "fmt"

func main() {
	usuario := map[string]string{ // dentro do colchete tipo das chaves, fora é o tipo do dado
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["time"] = map[string]string{
		"nome": "trikas",
	}
	fmt.Println(usuario2)

}
