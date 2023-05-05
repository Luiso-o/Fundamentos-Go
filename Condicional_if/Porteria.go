package main

import "fmt"

func main() {

	/*El portero de la discoteca decidirá dejar entrar a personas
	mayores de edad con la excepcion de que sea jueves*/

	//declaracionde variables
	edad := 18
	dia := "Jueves"

	/*En Go la condicion no va encerrada entre parentesis
	a no ser que declares mas de una condicion*/

	//si eres mayor de edad ó eres menor de edad y es jueves
	if (edad >= 18) || (dia == "Jueves" && edad >= 16) {
		fmt.Println("Puedes pasar! :D")

		//si tu edad es igual a 0
	} else if edad == 0 {
		fmt.Println("Tienes que darme una edad válida :/")

		//si eres menor de edad y no es jueves
	} else {
		fmt.Println("No puedes pasar >:( porque tienes", edad, "años")
	}

}
