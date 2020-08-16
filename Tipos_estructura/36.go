package main

import "fmt"

type User struct {
	Name string // Atributos
	Email string
}

func main() {
	
	// var cody User // Un objeto

	// cody.Name = "Cody"
	// cody.Email = "beppo@beppo.com"

	// cody := User { Name: "Cody", Email: "beppo@beppo.com" }

	Name := "Cody"
	Email := "beppo@beppo.com"

	cody := User { Name, Email }

	fmt.Println(cody)

}