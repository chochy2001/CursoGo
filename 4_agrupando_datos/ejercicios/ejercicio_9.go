package main

import "fmt"

func main() {
	//Usando el mapa anterior, agrega un nuevo contacto con el nombre
	//"Juan" y los teléfonos "12345678" y "87654321".
	//Borra el contacto de Juan anterior antes de agregar el nuevo contacto
	// Imprime el mapa actualizado.

	// ! Pidele al usuario el nombre del contacto a borrar y luego imprime el mapa actualizado.

	// ? Pista, usa un slice para guardar los teléfonos
	mapa := map[string][]uint32{
		"Juan":  {123456789, 987654321, 123323333},
		"Pedro": {324324236, 234534111},
		"Luis":  {423384029, 823984022, 1231231231, 222222, 11111},
	}

	for clave, valor := range mapa {
		fmt.Println("Nombre: ", clave)
		for i, valor2 := range valor {
			fmt.Printf("%d) Telefono: %d\n", i+1, valor2)
		}
	}
	mapa["Juan"] = []uint32{12345678, 87654321}

	fmt.Println("\n\n")
	for clave, valor := range mapa {
		fmt.Println("Nombre: ", clave)
		for i, valor2 := range valor {
			fmt.Printf("%d) Telefono: %d\n", i+1, valor2)
		}
	}

	var nombreContacto string
	fmt.Println("Escribe el nombre del contacto a borrar: ")
	fmt.Scanf("%s", &nombreContacto)
	delete(mapa, nombreContacto)
	fmt.Println("Se eliminó el contacto: ", nombreContacto)

	fmt.Println("\n\n")
	for clave, valor := range mapa {
		fmt.Println("Nombre: ", clave)
		for i, valor2 := range valor {
			fmt.Printf("%d) Telefono: %d\n", i+1, valor2)
		}
	}





}
