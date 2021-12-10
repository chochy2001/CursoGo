package main

import "fmt"

func main() {
	//Crea un mapa con tus contactos que tenga nombre y teléfono
	//Para cada contacto, imprime el nombre y el teléfono
	//Imprime el indice y el valor del mapa
	// ! Guarda más de 2 numeros por persona

	// ? Pista, usa un slice para guardar los teléfonos
	mapa := map[string][]uint32{
		"Juan":  {123456789, 987654321, 123323333},
		"Pedro": {324324236, 234534111},
		"Luis":  {423384029, 823984022, 1231231231, 222222, 11111},
	}
	/*
		fmt.Println(mapa)
		for indice, valor := range mapa {
			fmt.Println(indice, valor)
		}
	*/

	for clave, valor := range mapa {
		fmt.Println("Nombre: ", clave)
		for i, valor2 := range valor {
			fmt.Printf("%d) Telefono: %d\n", i+1, valor2)
		}
	}

}
