package main

import "fmt"

func main(){
	var edad uint8
	var nombre string
	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s", &nombre)
	fmt.Printf("Hola %s, tu edad es %d\n", nombre, edad)

	if edad == 10 {
		fmt.Println("Edad es igual a 10")
	} else if edad == 11 {
		fmt.Println("Edad es igual a 11")
	} else if edad == 12 {
		fmt.Println("Edad es igual a 12")
	} else if edad == 13 {
		fmt.Println("Edad es igual a 13")
	} else if edad == 14 {
		fmt.Println("Edad es igual a 14")
	} else if edad == 15 {
		fmt.Println("Edad es igual a 15")
	} else {
		fmt.Println("Edad es diferente a 10 o 15")
	}

/*
	// Switch
	switch edad {
	case 10:
		fmt.Println("Edad es igual a 10")
	case 11:
		fmt.Println("Edad es igual a 11")
	case 12:
		fmt.Println("Edad es igual a 12")
	case 13:
		fmt.Println("Edad es igual a 13")
	case 14:
		fmt.Println("Edad es igual a 14")
	case 15:
		fmt.Println("Edad es igual a 15")
	default:
		fmt.Println("Edad no esta entre 10 o 15")
	}

 */

	switch nombre {
	case "Juan":
		fmt.Println("Nombre es Juan")
	case "Pedro":
		fmt.Println("Nombre es Pedro")
	case "Maria":
		fmt.Println("Nombre es Maria")
	case "Jorge":
		fmt.Println("Nombre es Jorge")
	default:
		fmt.Println("Nombre no esta entre Juan, Pedro, Maria o Jorge")
	}
/*
	if nombre == "Juan" {
		fmt.Println("Nombre es Juan")
	} else if nombre == "Pedro" {
		fmt.Println("Nombre es Pedro")
	} else if nombre == "Maria" {
		fmt.Println("Nombre es Maria")
	} else if nombre == "Jorge" {
		fmt.Println("Nombre es Jorge")
	} else {
		fmt.Println("Nombre no esta entre Juan, Pedro, Maria o Jorge")
	}

 */









}
