package main

import "fmt"

func main() {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//ATRIBUIÇÂO

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONES
	//igual js

	// OPERADORES LÓGICOS
	//igual JS

	// OPERADORES UNÁRIOS
	//igual js

	// OPERADOR TERNÁRIO
	// não há, tem que fazer o if raiz
	numero := 5
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
