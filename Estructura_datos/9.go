package main

import "fmt"

func main() {
	
	usuarios := make(map[string] []int)

	usuarios["Saul"] = []int { 10, 9 , 8, 10}

	fmt.Println(usuarios)
}