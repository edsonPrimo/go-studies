package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	fmt.Println(posicao)

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	res := fibonacci(12)
	fmt.Println(res)
}
