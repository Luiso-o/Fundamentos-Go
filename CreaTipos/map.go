package main

import "fmt"

func main() {

	//MAP
	//Realiza un MAP con tres notas.
	notasAlumnos := map[string]int{

		"Francisco": 8,
		"Miriam":    5,
		"Rosa":      3,
	}

	fmt.Println(notasAlumnos)

	fmt.Println(notasAlumnos["Juan"]) //si el espacio no existe la consola devuelve un 0

}
