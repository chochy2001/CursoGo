package main

import "fmt"

func main() {
	mapa := map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}
	fmt.Println(mapa)
	mapa["cuatro"] = 4
	fmt.Println(mapa)

	// * delete(nombreMapa, nombreClave)
	// delete (mapa, "cuatro")
	//delete(mapa, "cuatro")
	delete(mapa, "dos")
	fmt.Println(mapa)

	if valor, ok := mapa["cinco"]; ok {
		fmt.Println(valor)
		delete(mapa, "cinco")
	} else {
		fmt.Println("No existe la clave cinco")
	}
	fmt.Println(mapa)

}
