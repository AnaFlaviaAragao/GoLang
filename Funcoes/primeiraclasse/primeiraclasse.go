package main

import "fmt"

// criando uma variavel que dentro tem uma funcao anonima
var soma = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))

}
