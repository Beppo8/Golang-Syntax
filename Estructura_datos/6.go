package main

import "fmt"

func main()  {
	
	meses := []string { "Enero", "Febrero", "Marzo"}

	// Un puntero
	// Una longitud
	// Una capacidad

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)

	meses = append(meses, "Octubre") // Si la estructura se encuentra en su capacidad maxima se genera un nuevo arreglo

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)
}