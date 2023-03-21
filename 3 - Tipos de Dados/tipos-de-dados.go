package main

import (
	"errors"
	"fmt"
)

func main() {
	const numero int = 1000000000000
	fmt.Println(numero)

	//alias
	// INT32 = rune
	var numero3 rune = 123456
	fmt.Println((numero3))

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 123123123123123.44
	fmt.Println(numeroReal1, numeroReal2)

	//STRINGS

	var str string = "texto"
	str2 := "texto2"

	fmt.Println(str, str2)

	char := 'A'
	fmt.Println(char)

	var error error = errors.New("deu ruim carai")
	fmt.Println(error)
}
