package main

import "fmt"

func main() {

	var numero = 1
	
	for ok := true; ok = numero < 10 { // Do While
		
		fmt.Println(numero)
		numero ++

	}

}