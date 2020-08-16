package main

import "fmt"
import "reflect"

func main() {
	 
	// var curso string = "Curso Profesional de Go!" // 1
	// var curso = "Curso profesinal de Go!" // 2
	curso := "Curso profesional de Go" // Listado

	longitud := len(curso) 

	fmt.Println(longitud)

	primerCaracter := curso[0] // Char -> uint8
	ultimoCaracter := curso[ len(curso) - 1]

	fmt.Println(primerCaracter)
	fmt.Println(reflect.TypeOf(primerCaracter))

	fmt.Println("%c\n", primerCaracter)
	fmt.Println("%c\n", ultimoCaracter)

}