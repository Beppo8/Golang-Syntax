package main 

import "fmt"

func main()  {
	
	var nombre string
	var edad int
	var altura float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre)

	fmt.Printf("Hola %s con una edad %d y una altura de %.2f\n", nombre, edad, altura)

}