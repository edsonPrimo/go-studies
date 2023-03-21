package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.idade = 21
	u.nome = "Carlin"
	fmt.Println(u)

	endereco1 := endereco{"calle 49", 1218}

	u2 := usuario{"Davi", 21, endereco1}
	fmt.Println(u2)

	u3 := usuario{nome: "Jorge"}
	fmt.Println(u3)

	u4 := usuario{"Samila", 12, endereco{"calle12", 123}}
	fmt.Println(u4)

}
