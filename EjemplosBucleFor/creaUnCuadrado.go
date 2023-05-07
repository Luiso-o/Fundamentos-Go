package main

import (
	"fmt"
)

func main() {
	//Dibuixa un cuadrat amb un únic * i que tingui 10 unitats d'ample i 10 d'alt.
	C := "*"
	for i := 0; i < 10; i++ { //ejecuta el for que tiene dentro 10 veces
		//fmt.Println(C)
		for j := 0; j < 10; j++ {
			fmt.Print(C + " ") //imprimirá una linea de 10 *
		}
		fmt.Println("") //hace un salto de linea
	}
}
