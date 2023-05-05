package main

import "fmt"

func main() {

	//Crea un secuencia de numeros del 0 al 20 de dos en dos

	//Estructura for con un parametro

	i := 0 //declaramos una variable fuera
	for i <= 20 {
		fmt.Print(i)

		//para imprimir los numeros separados por una coma y un punto final
		if i < 20 {
			fmt.Print(", ")
		} else {
			fmt.Print(".")
		}
		i += 2
	}

	fmt.Println("\n----------------------------------------")

	//Estructura for con dos parametros

	j := 0
	for {
		if j <= 20 {
			fmt.Print(j)
			j += 2

			//para imprimir los numeros separados por una coma y un punto final
			if j < 20 {
				fmt.Print(", ")
			} else {
				fmt.Print(".")
			}

		} else {
			break
		}
	}

	fmt.Println("\n----------------------------------------")

	//Estructura for con tres parametros

	for i := 0; i <= 20; i += 2 {
		fmt.Print(i)

		//para imprimir los numeros separados por una coma y un punto final
		if i < 20 {
			fmt.Print(", ")
		} else {
			fmt.Print(".")
		}
	}

}
