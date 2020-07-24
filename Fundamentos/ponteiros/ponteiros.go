package main

import "fmt"

func main() {
	//var i int = nil um inteiro não pode ter valor nil, seria 0
	i := 1

	
	var p *int = nil

	// Pegando o endereço da variável
	p = &i // pegue o endereço de i e atribua a p

	// *p 'acesso ao  valor' do qual aquele ponteiro aponta [valor]
	// p 'tenho o endereço' que está armazenado 'dentro do ponteiro p' [endereço]

	//Go não tem aritmetica de ponteiros
	*p++ // desreferenciando (pegando o valor)
	
	fmt.Println(i, *p)
	i++

	fmt.Println(p, *p, i, &i)

}
