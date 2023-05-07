package main

import (
	"fmt"
)

type alumno struct {
	nombre   string
	apellido string
}
type matriculado struct {
	alumno
	asignatura string
}

// Exemplo de metodo
func (a alumno) presentarse() { //De esta manera agregamos este metodo al struct alumno
	fmt.Println("Mi nombre es ", a.nombre, a.apellido, " y estoy aprendiendo mucho :D.")
}
func (m matriculado) presentarse() { //De esta manera agregamos este metodo al struct matriculado
	fmt.Println("Mi nombre es ", m.nombre, m.apellido, " y estoy cursando ", m.asignatura)
}

type usuario interface {
	presentarse()
}

func aula(u usuario) {
	fmt.Println("Estoy en el aula", u)
}
func main() {
	al1 := matriculado{
		alumno: alumno{
			nombre:   "Jon",
			apellido: "Ruiz",
		},
		asignatura: "Go",
	}

	pr1 := alumno{
		nombre:   "Joan",
		apellido: "Riera",
	}

	fmt.Println(al1)  //{{Jon Ruiz} Go}
	al1.presentarse() //Mi nombre es Jon Ruiz y estoy apriendo Go
	pr1.presentarse() //Mi nombre es Joan Riera  y estoy aprendiendo mucho.

	aula(pr1) //Estoy detro del aula {Joan Riera}
	aula(al1) //Estoy detro del aula  {{Jon Ruiz} Go}
}
