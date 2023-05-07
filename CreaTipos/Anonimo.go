package main

import "fmt"

func main() {
	/*A continuacion realiza 2 structs anonimos sobre articulos sobre
	articulos de supermecado donde tendrás que definir los datos nombre,
	unidades y pvp*/

	//struct anónimo (sin nombre especifico)
	producte1 := struct {
		nom          string
		unitats, pvp int
	}{
		nom:     "patates fregides",
		unitats: 63,
		pvp:     2,
	}
	producte2 := struct {
		nom          string
		unitats, pvp int
	}{
		nom:     "Coca cola",
		unitats: 13,
		pvp:     1,
	}

	fmt.Println("Producto 1 = ", producte1.nom, "\nUnidades = ", producte1.unitats, "\npvp = ", producte1.pvp)
	fmt.Println("")
	fmt.Println("Producto 2 = ", producte2.nom, "\nUnidades = ", producte2.unitats, "\npvp = ", producte2.pvp)

}
