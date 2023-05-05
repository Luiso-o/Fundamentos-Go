package main

import (
	"fmt"
)

var ciudad string = "barcelona"

func main() {

	cadena := fmt.Sprint("La ciudad de ", ciudad, " fue construida en Catalunya ")
	anio := 1800

	//cuando las variables son del mismo
	//tipo pueden concatenarse por comas con el signo de +
	fmt.Print(cadena + "desde 1800.\n")

	//se recomienda usar comas para concatenar variables de diferente tipo
	fmt.Print(cadena, "desde ", anio)

}
