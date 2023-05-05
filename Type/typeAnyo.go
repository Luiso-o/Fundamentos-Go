package main

import "fmt"

//tipo personalizado año con el tipo int subyacente
type anyo int

var anyoActual anyo

func main() {

	anyoActual = 2023
	fmt.Println(anyoActual)      //2023
	fmt.Printf("%T", anyoActual) //main.anyo

}
