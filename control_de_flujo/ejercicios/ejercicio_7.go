package main

import "fmt"

func main() {
	//Escribe un programa donde se use switch para imprimir el nombre de la semana en ingles
	//Que el usuario lo escriba en español y lo imprimas en ingles.
	var dia string
	fmt.Println("Escribe 'salir' para salir")

	for dia != "salir" {

		fmt.Println("Ingresa el dia de la semana en español")
		fmt.Scanf("%s", &dia)

		switch dia {
		case "lunes", "Lunes", "LUNES":
			fmt.Println("Monday")
		case "martes", "Martes", "MARTES":
			fmt.Println("Tuesday")
		case "miercoles", "Miercoles", "MIERCOLES":
			fmt.Println("Wednesday")
		case "jueves", "Jueves", "JUEVES":
			fmt.Println("Thursday")
		case "viernes", "Viernes", "VIERNES":
			fmt.Println("Friday")
		case "sabado", "Sabado", "SABADO":
			fmt.Println("Saturday")
		case "domingo", "Domingo", "DOMINGO":
			fmt.Println("Sunday")
		default:
			fmt.Println("No es un dia de la semana")
			break
		}
	}

}
