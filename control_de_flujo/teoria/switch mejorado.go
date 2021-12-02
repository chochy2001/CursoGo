package main

import "fmt"

func main() {
	var edad uint8
	var nombre string
	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s", &nombre)

	switch edad {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("Eres un bebé")
	case 10, 11, 12, 13, 14, 15, 16, 17:
		fmt.Println("Eres un niño")
	case 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29:
		fmt.Println("Eres un adulto joven")
	default:
		fmt.Println("Eres un adulto")
	}

	switch nombre {
	case "Juan", "juan":
		fmt.Println("Nombre es Juan")
	case "Pedro", "pedro":
		fmt.Println("Nombre es Pedro")
	case "Maria", "maria":
		fmt.Println("Nombre es Maria")
	case "Jorge", "jorge":
		fmt.Println("Nombre es Jorge")
	default:
		fmt.Println("Nombre no esta entre Juan, Pedro, Maria o Jorge")
	}

	switch {
	case edad < 18 && nombre == "Juan":
		fmt.Println("Eres menor de edad Juan")
	case edad < 18 && nombre == "Pedro":
		fmt.Println("Eres menor de edad Pedro")
	case edad>=18 && nombre == "Maria":
		fmt.Println("Eres mayor de edad Maria")
	default:
		fmt.Println("No se puede determinar")
	}

}
