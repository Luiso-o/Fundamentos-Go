package main

import (
	"fmt"
)

var edat int
var ciutat string = "Barcelona"

func main() {

	//Imprimiento solamente la variable
	fmt.Println(edat)
	fmt.Println(ciutat)

	//Agregando un mensaje concatenando la variable
	fmt.Printf("El valor de edat es: %v\n", edat)

	//Se concatena con una coma cuando las variables son diferentes
	fmt.Printf("El valor de ciutat es: %v", ciutat)
}
