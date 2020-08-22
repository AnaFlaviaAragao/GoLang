package main

import "fmt"

func main() {

	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)  Não é possivel
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	// O copy primeiro parametro é ao que recebe o valor copiado, e o segundo paramentro é a fonte
	// O copy não espande o tamanho do slice, ou seja, o tamanho vai ser coerente ao primeiro parametro
	// copia de uma fonte para outra
	// No copy o segundo parametro não pode ser um array como fonte de dados
	// No copy obrigatoriamente tem que ser um slice/ sring
	copy(slice2, slice1)
	fmt.Println(slice2)

}
