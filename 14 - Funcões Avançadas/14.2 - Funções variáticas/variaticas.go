package main

import "fmt"

func soma(numeros ...int) (total int) {
	for _, numero := range numeros {
		total += numero
	}
	return
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma := soma(1, 2, 3, 4, 532423, 567, 7, 8, 9)
	fmt.Println(soma)

	escrever("Texto doido", 123123, 123, 123, 123, 123)
}
