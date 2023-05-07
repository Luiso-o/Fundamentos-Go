package main

import "fmt"

func main() {

	carro := []int{23452, 72622, 5433, 9273, 62533, 62353, 15625, 25162, 41312, 62762, 62532, 52426}
	total := 0

	/*Teniendo en cuenta la cantidad de datos mostrados, correspondiendo
	a la facturacion de uan tienda de cada mes del año, tienes que
	mostrar el total facturado del primer trimestre*/

	fmt.Println(carro[:3])

	for _, valor := range carro[:3] { //:3 recorremos y nos detemos al llegar al indice 3
		total += valor
	}

	fmt.Println("El total facturado este trimestre es de :  ", total, "euros")
}

//Cuando queremos cancelar una variable usamos un "_", asi podemos seguir trabajando
//y revisar la situacion de la variable despues.
