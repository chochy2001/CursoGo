package main

import "fmt"

func main() {
	// * IF
	//if condición {ejecuta este código}

	/*
		if 10 < 5{//si la condición es verdadera, entonces ejecuta el código
			fmt.Println("10 es menor que 5")
		}
		// * Operador !
		// ! negación true -> false    false -> true
		a := 10
		b := 5

		if !(a<b) {//para que entren en el if, la condición debe ser verdadera
			fmt.Println("a es menor que b")
		}


		if true{
			fmt.Println("Verdadero Normal")
		}
		if false{
			fmt.Println("Falso Normal")
		}
		if !true{
			fmt.Println("Verdadero Negado")
		}
		if !false{
			fmt.Println("Falso Negado")
		}
		if 2==2{
			fmt.Println("2 es igual a 2")
		}
		if 2!=2{
			fmt.Println("2 es diferente a 2")
		}
		if edad:= 10; edad > 18 {
			fmt.Println("Mayor de edad")
		}

		// * Operador &&
		// && and
		fmt.Println("&&")
		if edad := 10; edad > 18 && edad < 30 {
			fmt.Println("Eres Joven")
		}
		//&& lo que hace es que si la condición es verdadera, entonces ejecuta el código
		//Si tenemos dos sentencias, ambas tienen que ser verdaderas para que entre en el if





		// * Operador ||
		// || or
		fmt.Println("||")
		if edad := 10; edad > 18 || edad < 30 {
			fmt.Println("Eres Joven")
		}
		//|| lo que hace es que si la condición es verdadera, entonces ejecuta el código
		//Con que una condicion sea verdadera, ya entra en el if

		edad := 18
		if edad > 18 && edad < 30 {
			fmt.Println("Puedes pasar")
		}else if edad > 30{
			fmt.Println("No Puedes pasar, ya estas viejo")
		}else if edad < 18{
			fmt.Println("No Puedes pasar, eres muy joven")
		}else{
			fmt.Println("No se que eres")
		}
		// * else if
		// else if es una condición que puede ser verdadera o falsa
		// else if puede tener una condicion
		/*if condicion{

		}else if condicion{

		}else if condicion{

		}else{

		}






		// * else
		//else es una condición que puede ser verdadera o falsa
		// else no tiene una condicion
		if condicion{

		}else{

		}
	*/

	//Programa que con base en la edad del usuario te dice si es mayor o menor de edad
	var edad uint8
	fmt.Println("Ingresa tu edad")
	fmt.Scanf("%d", &edad)
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
