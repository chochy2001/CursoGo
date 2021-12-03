package main

import "fmt"

func main() {
	//Crea un programa que si eres mayor de edad,
	//te diga "Eres mayor de edad" y sino eres mayor de edad , te diga "Eres menor de edad"
	var edad uint8

	fmt.Println("Para salir del programa presiona el 0\n")

	/*
	for{
		fmt.Println("Escribe tu edad: ")
		fmt.Scanf("%d", &edad)

		if edad >= 18 {
			fmt.Println("Eres mayor de edad\n")
		} else {
			fmt.Println("Eres menor de edad\n")
		}
		if edad == 0 {
			break
		}
	}
	 */

	// si tu edad es igual a 17, "si me das dinero eres mayor de edad"
	for{
		fmt.Println("Escribe tu edad: ")
		fmt.Scanf("%d", &edad)

		if edad >= 18 {
			fmt.Println("Eres mayor de edad\n")
		}else if edad == 17 {
			fmt.Println("Si me das dinero eres mayor de edad\n")
		} else {
			fmt.Println("Eres menor de edad\n")
		}

		if edad == 0 {
			break
		}
	}

}
