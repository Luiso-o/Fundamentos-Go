package main //Declaracion del archivo actual

import "fmt" //paquete predeterminado para imprimir

func main() { //funcion principal

	//declaracion de variable - nombre de la variable - tipo pode variable
	var a int     //entero
	var b string  //cadena de texto
	var c float64 //decimal
	var d bool    //opcion binaria

	//fmt.printf = se utiliza para imprimir la variable
	// %T = Impresion del tipo variable
	// %v = Impresion del valor de la variable
	// &q = Formato escapado

	fmt.Printf("var a %T = %+v\n", a, a) //salida var a int = 0

	fmt.Printf("var b %T= %q\n", b, b) //salida var b string = ""

	fmt.Printf("var c %T = %+v\n", c, c) //salida var c float64 = 0

	fmt.Printf("var d %T = %+v\n", d, d) //salida var d bool = false
	
	//En Go no se usa el ;

}
