package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha")

	fmt.Println("linha.")
	fmt.Println("linha1.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de X é ", x)

	fmt.Printf("O valor de x é %.2f ", x)

	a := 1
	b := 1.99999
	c := false
	d := "Opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// %v geral, porém não funciona para todos. 
	fmt.Print("\n %v %v %v %v", a, b, c, d)
}
