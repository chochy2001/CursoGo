package main

import "fmt"

func main() {
	//Escribe un programa en donde se use el switch sin ninguna condición.
	for {
		var edad uint8
		var nombre string
		fmt.Println("Ingrese su edad: ")
		fmt.Scanf("%d", &edad)
		fmt.Println("Ingresa tu nombre")
		fmt.Scanf("%s", &nombre)

		switch {
		case edad < 18 && nombre == "Juan":
			fmt.Println("Eres menor de edad Juan")
		case edad < 18 || nombre == "Pedro":
			fmt.Println("Eres menor de edad Pedro")
		case edad >= 18 && nombre == "Maria":
			fmt.Println("Eres mayor de edad Maria")
		case edad == 17:
			fmt.Println("tu edad es igual a 17")
		case !(edad == 17) && nombre == "Jorge":
			fmt.Println("tu edad es igual a 17 y tu nombre es Jorge")
		case edad == 17 || nombre == "Jorge":
			fmt.Println("tu edad es igual a 17 o tu nombre es Jorge")
		default:
			fmt.Println("No se puede determinar")
		}

		if edad == 0 && nombre == "salir" {
			break
		}

	}

}
