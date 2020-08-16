package main

import "fmt"

// Retornar funciones

type Operacion func(balance, cantidad int) int

func crearOperacion(tipo string) Operacion {

	if tipo == "suma" {
		return func(balance, cantidad int) int { return balance + cantidad }
	} else if tipo == "resta" {
		return func(balance, cantidad int) int { return balance - cantidad }
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}

}

func main() {

	suma := crearOperacion("suma")

	resultado := suma(40, 50)

	fmt.Println("El resultado es:", resultado)

}