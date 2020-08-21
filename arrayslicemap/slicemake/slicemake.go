package main

import "fmt"

func main() {
	//:= --> criar, iniciar e definir
	//criar slice
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//reatribuindo pois S ja existe
	// terceiro parametro de array interno ou seja, capacidade.
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
