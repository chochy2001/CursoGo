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

	for clave, valor := range mapa {
		fmt.Println(clave, valor)
	}

	slice := []string{"uno", "dos", "tres", "cuatro"}
	for indice, valor := range slice {
		fmt.Println(indice, valor)
	}

}
