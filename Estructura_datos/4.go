package main

import "fmt"

func main()  {
	
	// var numeros[5] int

	numeros := [...] int { 100, 200 , 300, 400}

	numeros[0] = 100

	fmt.Println(numeros)

	// Mas arrays

	monedas := [...]string {0: "Dolar Canadiense", 1: "Peso Mexicano", 2:"Dolar", 5:"Euro"}

	subMonedas := monedas[0:3] // Slice

	

}