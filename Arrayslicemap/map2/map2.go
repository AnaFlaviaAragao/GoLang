package main

import "fmt"

func main() {
	// Inicializar na propria declaração
	funcESalarios := map[string]float64{
		"José João":   11325.45,
		"Gabriel":     15456.78,
		"Pedro Lucas": 1200.0,
	}
	// Adicionar
	funcESalarios["Rafael Filho"] = 1350.0
	// Remover
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}

	fmt.Println("\n")

	// Imprimir a chave
	for salario := range funcESalarios {
		fmt.Println(salario)
	}

}
