package main

import "fmt"

// Relacion OneToOne

type User struct {
	Name string
	Email string
	Active bool
}

type Student struct {
	User User
	Id string
}

func main() {
	
	eduardo := User{ Name: "Eduardo", Email: "eduardo@codigofacilito.com", Active: true }

	uriel := User{ Name: "Uriel", Email: "uriel@codigofacilito.com", Active: true }

	eduardoStudent := Student { User: eduardo, Id: "1f1" }

	fmt.Println(uriel)
	fmt.Println(eduardoStudent.User.Email)

}