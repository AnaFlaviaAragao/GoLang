package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador contar!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros { /// Ignora os indices
		//fmt.Printf("%d\n", num)
		fmt.Println(num)
	}

	for num := range numeros { //Pega os indices
		//fmt.Printf("%d\n", num)
		fmt.Println(num)
	}
}
