package main

import "fmt"

// método para inicializar qualquer tipo de proceso antes da inicialização
//é executada quando é executada
func init() {
	fmt.Println("Inicializando...")
}
func main() {
	fmt.Println("Main...")

}
