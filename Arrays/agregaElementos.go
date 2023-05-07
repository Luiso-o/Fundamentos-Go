package main

import "fmt"

//Slice
func main() {

	carro := []string{"Zanahorias", "Pimientos", "Agua", "Queso"}

	//Agrega al carro de la compra Aceite de Oliva
	//nombre de la variable = append (variable, Elementos que deseas agregar)
	carro = append(carro, "Aceite de Oliva")

	//Imprimimos
	fmt.Println(carro)

}
