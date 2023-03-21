package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println("Incrementando", i)
		i++
	}

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("Incrementando", j)
	}

	// nomes := [3]string{"João", "Paulo", "Rebeca"}
	nomes2 := []string{"CARLINHOS", "SEVERINO"}

	for _, nome := range nomes2 { // primeiro parametro: indice, segundo: item do array
		time.Sleep(time.Second)
		fmt.Println("nome", nome)
	}

	for indice, letra := range "ASDKAPSOKDOASKDOAS" { // primeiro parametro: indice, segundo: item do array
		time.Sleep(time.Second)
		fmt.Println("indice", indice)
		fmt.Println("letra", string(letra))
	}

	user := map[string]string{
		"nome":  "paulo",
		"idade": "12",
	}

	for _, valor := range user {
		fmt.Println("valor", valor)
	}

}
