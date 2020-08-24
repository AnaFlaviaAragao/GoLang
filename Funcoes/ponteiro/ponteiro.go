package main

import "fmt"

func inc1(n int) {
	n++ // n = 1 + n
}

//revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta

	*n++ //não tem aritmetica de ponteiro, podemos referenciar e ter acesso ao valor

}

func main() {
	n := 1
	inc1(n) // por valor
	fmt.Println(n)

	// revisão & usada para o obter o endereço da variavel
	inc2(&n) //por referencia
	fmt.Println(n)
}
