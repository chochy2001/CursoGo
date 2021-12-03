// * Escribe para que sirven las distintas expresiones
// == != < > <= >=
// * Usa las expresiones de arriba para comprobar que funcionan e imprime su resultado

//! Pidele la edad al usuario y que al momento de ingresarla
//! uses los operadores de arriba e imprimas su resultado comparandolo con 18

// == Igual
// != Diferente
// < Menor que
// > Mayor que
// <= Menor o igual que
// >= Mayor o igual que

package main

import "fmt"

func main() {
	var edad uint8

	/*
		fmt.Println(10 == 18)
		fmt.Println(10 != 18)
		fmt.Println(10 < 18)
		fmt.Println(10 > 18)
		fmt.Println(10 <= 18)
		fmt.Println(10 >= 18)
	*/

	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Println("Su edad es: ", edad)

	fmt.Printf("%d == 18: %t\n", edad, edad == 18)
	fmt.Printf("%d != 18: %t\n", edad, edad != 18)
	fmt.Printf("%d < 18: %t\n", edad, edad < 18)
	fmt.Printf("%d <= 18: %t\n", edad, edad <= 18)
	fmt.Printf("%d > 18: %t\n", edad, edad > 18)
	fmt.Printf("%d >= 18: %t\n", edad, edad >= 18)

}
