package main

import "fmt"

// Funciones recursivas

func factorial(numero int) int {

	if numero == 1 {
		return 1
	}

	return factorial(numero - 1) * numero

}

func main() {
	
	resultado := factorial(5)
	fmt.Println("El factorial es:", resultado)

}
