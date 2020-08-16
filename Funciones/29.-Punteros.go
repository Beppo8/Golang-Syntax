package main

import "fmt"

func modificarValor(parametro *string) {
	*parametro = "Cambio de valor"
}

func main() {
	
	var curso = "Curso profesional de Go!"

	fmt.Println("Antes de la funcion:", curso)

	modificarValor(&curso) // Referencia

	fmt.Println("Despues de la funcion:", curso)

}