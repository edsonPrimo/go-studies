package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 2:
		return "Domingo"
	default:
		return "Número Inválido"
	}
}

func main() {
	dia := diaDaSemana(1)
	fmt.Println(dia)
}
