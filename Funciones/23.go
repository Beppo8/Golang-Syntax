package main

import "fmt"
 
// Funciones anonimas

func main() {
	
	miFuncion := func (username string) string {
		message := fmt.Sprintf("Hola %s, te saludamos desde una funcion sin nombre", username)

		return message
	}

	mensajeFinal := miFuncion("Cody 1")

	fmt.Println(mensajeFinal)

}