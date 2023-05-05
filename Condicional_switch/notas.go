package main

import "fmt"

func main() {

	nota := 8

	switch nota { //la condicion no va entre parentesis

	case 1, 2, 3, 4: //los casos pueden ser multiples y separados por comas
		fmt.Println("has suspendido, esfuerzate más :'(")

	case 5:
		fmt.Println("has aprobado, sigue practicando :)")

	case 6, 7, 8:
		fmt.Println("Has sacado un notable, sigue así :D")

	case 9, 10:
		fmt.Println("has sacado un excelente, muchas felicidades! :'D")
		//el brack no es necesario
	}
}
