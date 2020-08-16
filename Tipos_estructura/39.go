package main

import "fmt"

// Relacion OneToMany

type Curso struct {
	Titulo string
	Videos []Video
}

type Video struct {
	Titulo string
	Curso Curso
}

func main() {
	
	video1 := Video { Titulo: "1.-Introduccion" }
	video2 := Video { Titulo: "2.-Introduccion" }

	videos := []Video { video1, video2 }

	curso := Curso { Titulo: "Curso profesional de Go!", Videos: videos }
	
	video1.Curso = curso
	video2.Curso = curso

	fmt.Println(video1)
	fmt.Println(video2)

	fmt.Println(video1.Curso.Titulo)

	for _, video := range curso.Videos {
		fmt.Println(video.Titulo)
	}

}