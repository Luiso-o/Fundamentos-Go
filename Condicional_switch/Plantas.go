/*
Realizar una estructura condicional que contemple
las opciones del ascensor de un edificio:
1. Planta 1 y 3 = Articulos del hogar.
2. Planta 2 = Moda Infantil
3. Planta 4 = Area Jugueteria
4. Default error
*/
package main

import "fmt"

func main() {

	planta := 3

	switch planta {
	//caso planta 1 y panta 3 articulos del hogar
	case 1, 3:

		//divido las plantas del hogar en diferentes zonas
		if planta == 1 {
			fmt.Println("Articulos del hogar: Dormitorio, cocina, sala")
		} else {
			fmt.Println("Articulos del hogar: Baño, Jardin, trastero")
		}

	//caso planta 2 moda infantil
	case 2:
		fmt.Println("Moda infantil")

	//caso planta 4 area jugueteria
	case 4:
		fmt.Println("Area de jugueteria")

	//si no tenemos un edeficio con este numero
	default:
		fmt.Println("No existe esta planta, escoge una planta entre 1 i 4.")

	}

}
