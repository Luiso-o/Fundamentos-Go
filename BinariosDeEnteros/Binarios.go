package main

import (
	"fmt"
)

func main() {

	//variable de tipo int
	a := 42

	//d% = valor decimal de la variable
	//b% = valor binario de la variable
	fmt.Printf("%d\t%b\n", a, a) //42  101010

	//Usamos el operador <<

	b := a << 1

	/*El operador << realiza una operación de desplazamiento a
	la izquierda en un número binario. En este caso, se desplaza
	el valor binario de a una posición a la izquierda, lo que
	 equivale a multiplicar el valor de a por 2. Por lo tanto,
	 si a tiene el valor de 42 en decimal, su valor binario es
	  101010. Al desplazar 101010 una posición a la izquierda,
	  se obtiene 1010100, que es el valor binario de 84.*/

	fmt.Printf("%d\t%b\n", b, b) //84 1010100

}
