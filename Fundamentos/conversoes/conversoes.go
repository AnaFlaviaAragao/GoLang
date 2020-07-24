package main

import (
	"fmt"
	c "strconv"
)

func main() {
	/*x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)*/

	//cuidado
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + c.Itoa(123))

	//string para int
	num, _ := c.Atoi("123")
	fmt.Println(num - 122)

	//string para boolean 1 ou true
	b, _ := c.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
