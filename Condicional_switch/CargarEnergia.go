/* realizar una estructura condicional switch que contemple lo siguiente
1. que tenga energia al 100% o tenga un potenciador, mostrando el mensaje "Energia cargada
2. si la energia es menor a 100%, aunmentar la energia 20 unidades y mostrar un mensaje de la energia actual
3. un default, mostrando un error"*/

package main

import "fmt"

func main() {

	energia := 50
	booster := false

	//en Go switch no usa condicional

	switch { //si la energia está cargada al 100% o el potenciador está activado
	case energia == 100, booster == true:
		fmt.Println("Energia cargada al 100%")

	case energia < 100:
		fmt.Print("Energia al ", energia, "% Cargando...\n")

		energia += 20

		if energia > 100 { //si la carga supera el 100% la energia se quedará al 100%
			energia = 100
		}

		fmt.Println("Ahora mismo tines un", energia, "% de energia")

	default:
		fmt.Println("Error X_X")
	}

}
