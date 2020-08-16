package main

import "fmt"

// Funciones con argumentos

type Operacion func(balance, cantidad int) int

func suma(numero1, numero2 int) int { // Operacion
	return numero1 + numero2
}

func resta(numero1, numero2 int) int { // Tipo Operacion
	return numero1 - numero2
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {

	fmt.Println("Antes de la operacion")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es:",resultado)

	fmt.Println("Despues de la operacion")

}

func main() {
	
	ejecutarOperacion(resta, 100, 50)

}