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

// Ejemplo de metodos (funciones)
func (m matriculado) presentarse() { //De esta manera agregamos este metodo al struct matriculado
	fmt.Println("Mi nombre es ", m.nombre, m.apellido, " y estoy aprendiendo ", m.asignatura)
}
func main() {
	al1 := matriculado{
		alumno: alumno{
			nombre:   "Jon",
			apellido: "Ruiz",
		},
		asignatura: "Go",
	}
	fmt.Println(al1)  //{{Jon Ruiz} Go}
	al1.presentarse() //Mi nobre es Jon Ruiz y estoy aprendiendo Go
}
