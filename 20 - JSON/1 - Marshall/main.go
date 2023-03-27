package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	dog1 := cachorro{Idade: 5, Nome: "Nico", Raca: "Cocker"}
	fmt.Println(dog1)

	cachorroEmJSON, err := json.Marshal(dog1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	dog2 := map[string]string{
		"nome": "Bolota",
		"raca": "SRV",
	}

	cachorro2EmJSON, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

}
