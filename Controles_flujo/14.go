package main

import "fmt"

func main() {
	
	usuarios := make(map[string]string)

	usuarios["beppo"] = "beppo@codigofacilito.com"

	if correo, ok := usuarios["cody"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("no fue posible obtener el valor")
	}

}