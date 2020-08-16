package main

import "fmt"

func main()  {
	
	numeros := []int { 1, 2, 3, 4} // Referencia a un arreglo base.

	numeros = append(numeros, 5)

	nuevoSlice := numeros[0:5]

	numeros[0] = 100

	fmt.Println(numeros)
	fmt.Println(nuevoSlice)
}