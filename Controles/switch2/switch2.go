package main

import (
	"fmt"
	"time"
)

//Desafio
func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n <= 9:
		return "B"
	case n >= 5 && n <= 8:
		return "C"
	case n >= 3 && n <= 5:
		return "D"
	default:
		return "E"
	}
}

func main() {

	t := time.Now()
	switch { //switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde")
	default:
		fmt.Println("Boa noite")
	}

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.8))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(4.8))
}
