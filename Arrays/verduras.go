package main

import "fmt"

//Arrays en Go

func main() {

	//nombre del array/tendra el valor de/ un array variable de cadenas de texto.
	carro := []string{ //array variable
		"pastanagues", "pebrots", "aigua", "formatge"}

	carro[1] = "cogombre" //puedes sustiruir elementos en la lista

	fmt.Println(carro)

}
