package main

import (
	"fmt"
)

// struct persona
type persona struct {
	nombre string
	edad   int
}

// invoca al struct persona(similar a la herencia)
type ingenier struct {
	persona
	realitzarPlanol bool
}

func main() {

	//puedes usar un estruct dentro de otro struct
	//le damos valor a las variables
	ing1 := ingenier{
		persona: persona{
			nombre: "Ruben",
			edad:   26,
		},
		realitzarPlanol: true,
	}

	fmt.Println(ing1)
	//los atributos de persona pasan a ser parte del struct ingener
	fmt.Println(ing1.nombre, ing1.edad, ing1.realitzarPlanol)
}
