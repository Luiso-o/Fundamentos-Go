package main

import (
	"fmt"
)

func main() {

	//make (metodo para realizar un slice)
	//make = (tipo de slice, longitud, capacidad)
	x := make([]string, 9, 10)
	x[3] = "Joan"
	x[7] = "Silvia"
	fmt.Println(x)
	fmt.Println(len(x)) //len = longitud
	fmt.Println(cap(x)) //cap = capacidad

	x = append(x, "Cristina") //agregamos una nueva cadena
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, "Ana")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

/*x := make([]int, 9, 10)

x[2] = 3
x[6] = 7

fmt.Println(x)
fmt.Println(len(x))
fmt.Println(cap(x))

x = append(x, 10)
fmt.Println(x)
fmt.Println(len(x))
fmt.Println(cap(x))

//x[11] = 11 //Això donarà error
x = append(x, 11)
fmt.Println(x)
fmt.Println(len(x))
fmt.Println(cap(x)) //Això duplicarà el espai de memòria*/
