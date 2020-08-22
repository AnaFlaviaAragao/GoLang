package main

import "fmt"

//chave(não aceita repetição) valor

func main() {
	// Mapas devem ser inicializados
	//var aprovados map[int]string // chave(int) valor(string)

	aprovados := make(map[int]string)

	aprovados[893147983478367] = "Maria"
	aprovados[844987542728775] = "Pedro"
	aprovados[349842842737474] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF): %d\n", nome, cpf)
	}

	fmt.Println(aprovados[349842842737474])
	//deletar
	delete(aprovados, 349842842737474)
	fmt.Println(aprovados[349842842737474])

	// OBS. chaves iguais os valores seram sobescritos

}
