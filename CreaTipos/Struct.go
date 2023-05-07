package main

import (
	"fmt"
)

// struct persona(los atributos se declaran dentro)
type persona struct {
	nom  string
	edat int
}

func main() {

	//instanciar struct
	p1 := persona{
		nom:  "Gerard",
		edat: 22,
	}

	//instanciar struct
	p2 := persona{
		nom:  "Ona",
		edat: 18,
	}

	//imprimimos los datos
	//todo el objeto
	fmt.Println(p1)
	fmt.Println(p2)

	//tambien se puede acceder por atributos
	fmt.Println(p1.nom)
	fmt.Println(p2.nom)
}
