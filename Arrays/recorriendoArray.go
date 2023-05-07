package main

import "fmt"

//recorer un array
func main() {
	//Mira de recorrer i mostrar les dades atraves d'un for amb range.
	carro := []string{"Zanahorias", "Pimientos", "Agua", "Queso"}
	ID := 1

	//recorro usando un for con los parametros inice del array y valor del indice
	for index, valor := range carro {

		//index = {0 1 2 3} , valor = "pastanagues", "pebrots", "aigua", "formatge"
		fmt.Println("Item ", index, " => Producte: ", valor, "codigo =>", ID)
		ID++

	}
}
