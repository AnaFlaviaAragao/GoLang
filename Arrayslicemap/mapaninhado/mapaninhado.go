package main

import "fmt"

func main() {
	// Temos um map com chave estring e =como valor temos outro map
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gariela Silva": 1544.12,
			"Guga Perreira": 1554.23,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra["G"], "Gariela Silva")
	delete(funcsPorLetra, "P")

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {

		fmt.Println(letra, funcs)
	}

}
