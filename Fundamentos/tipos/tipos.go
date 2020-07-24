package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// interger numbers
	fmt.Println(1, 2, 10000)
	fmt.Println("Literal Inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... unit8[1B 8bits 0-256] uint16[2B] uint32[4B] uint64[8b]
	var b byte = 255
	fmt.Println("0 byte é", reflect.TypeOf(b))

	//Com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor proximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais ( float 32, float 64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal de 4.99", reflect.TypeOf(49.99))

	//boleab
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	sl := `Olá meu nome é Agatha`
	fmt.Println(sl + "!")
	fmt.Println("O tamanho da string é", len(sl))

	//Não tem char
	char := 'a'
	fmt.Println("O tamanho da string é", reflect.TypeOf(char))
	fmt.Println(char)

}
