package main

import "fmt"

func main() {
	numeros := []int{9, 3, 12, 7, 5, 2, 8}
	i, acumulador := 0, 0
	for {
		if i == len(numeros) {
			break
		}
		//condicio per esbrinar si els nombres son parells
		if numeros[i]%2 == 0 {
			acumulador += numeros[i]
			fmt.Println(numeros[i], "es un numero par.")
		}
		i++
		continue
	}

	fmt.Println("La suma de los numeros pares es: ", acumulador)

}
