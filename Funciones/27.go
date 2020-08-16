package main

import "fmt"

func modificarVariable(parametro string) {

	fmt.Pritnln("Valor actual:", parametro)
	parametro = "Nuevo curso de Go!"
	fmt.Pritnln("Nuevo valor:", parametro)

	fmt.Pritnln("%p", &parametro)

}

func main() {
	
	var curso = "Curso profesional de Go!"

	modificarVariable(curso)

	fmt.Pritnln(">>>", curso)
	fmt.Pritnln("p" &curso)

}