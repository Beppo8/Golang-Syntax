package main

import "fmt"

// Variadic function
func promedio(calificaciones ...int) int {

	var promedio, suma int

	for _, calificacion := range calificaciones {
		suma = suma + calificacion
	}

	promedio = suma / len(calificaciones)

	return 0

}

func main() {
	
	promedio(10, 8, 7, 10, 9)

	fmt.Println("Promedio:", resultado)

}