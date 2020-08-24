package main

import "fmt"

// defer atrasa, ser executada para ser atrasada
// usada para fechar algum tipo de recurso que se abriu

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!") // Atrasa a execução para o ultimo momento/ ou momento mais tardio dentro desse metodo
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
