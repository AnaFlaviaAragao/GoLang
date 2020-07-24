package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inerencia de tipo
	i += 3 // i = i + 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3

	fmt.Println(i)
	x, y := 1, 2 // criar e  inicializar
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
