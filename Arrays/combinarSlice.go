package main

import "fmt"

//juntar dos Slices
func main() {

	//Slice1
	trabajadores := []string{"Josep", "Cristina", "Helena", "Robert"}

	//Slice2
	nuevasIncorporaciones := []string{"Ivan", "Estel", "Aitana"}

	//El slice que quieres que predomine = append(slice1,slice2 más ...)
	trabajadores = append(trabajadores, nuevasIncorporaciones...)

	//imprimes el slice predominante
	fmt.Println(trabajadores)
}
