package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice - tamanho variavel
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array!  é um trecho de um array.
	// craia uma estrutura que aponta e não cria outro array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice mas internamente está apontando par ao mesmo array
	fmt.Println(a2, s3)

	// Podemos imaginar um slice como: tamanho e um ponteiro par aum elemento de um array

	s4 := s2[:1]
	fmt.Println(s2, s4)

}
