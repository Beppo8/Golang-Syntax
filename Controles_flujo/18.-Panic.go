package main

import "fmt"

func main() {
	
	var dividendo , divisor int

	fmt.Println("Ingresa un valor para el dividendo")
	fmt.Scanf("%d" &dividendo)

	fmt.Println("Ingresa un valor para el divisor")
	fmt.Scanf("%d" &divisor)

	if divisor == 0 {
		panic("No es posible dividir sobre 0")
	}

	resultado:= dividendo / divisor

	fmt.Println(resultado)

}