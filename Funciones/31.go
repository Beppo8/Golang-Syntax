package main

import "fmt"

// Funciones como valores

func factorial(numero int) int {

	if numero == 1 {
		return 1
	}

	return factorial(numero - 1) * numero

}

type customeFunction func(n int) int

func main() {
	
	// var miFuncion = factorial
	var miFuncion = customeFunction // nil

	if miFuncion == nil {
		miFuncion = factorial
	}

	resultado := miFuncion(3)

	fmt.Println(resultado)

}