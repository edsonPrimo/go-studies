package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println("Ola Mundo", texto)
	}("texticulo")
}
