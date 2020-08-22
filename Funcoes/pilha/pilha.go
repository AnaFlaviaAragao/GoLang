package main

import "runtime/debug"

//Printar a pilha de exec de um  programa
func f3() {
	debug.PrintStack()
}
func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
