package main

import "fmt"

func main() {
	
	var calificacion int

	fmt.Print("Ingresa una calificacion: ")
	fmt.Scanf("%d", &calificacion)

	/*
	switch {
	case calificacion == 10:
		fmt.Println("Felicidades obstuviste una calificacion perfecta!")
	case calificacion == 8 || calificacion == 9:
		fmt.Println("Aprobaste la materia")
	case calificacion == 7 || calificacion == 6:
		fmt.Println("Aprobaste la materia, pero necesitas estudiar mas.")
	case calificacion <= 5 && calificacion >= 0:
		fmt.Println("No aprobaste la materia")
	default:
		fmt.Println("Calificacion no valida")
	}
	*/

	switch calificacion {
		case 10:
			fmt.Println("Felicidades obstuviste una calificacion perfecta!")
		case 8, 9:
			fmt.Println("Aprobaste la materia")
		case 6, 7:
			fmt.Println("Aprobaste la materia, pero necesitas estudiar mas.")
		case 1, 2, 3, 4, 5:
			fmt.Println("No aprobaste la materia")
		default:
			fmt.Println("Calificacion no valida")
	}



}