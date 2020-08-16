package main

import "fmt"

func suma(numero1 , numero2 int) (int, string) {

	return numero1 + numero2, "Un mensaje desde la funcicon suma"

}

func main() {
	
	resultado, _ := suma(10, 20)	

	fmt.Println(resultado)

}