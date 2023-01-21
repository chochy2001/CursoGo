package main

import "fmt"

func main() {
	a := 10

	var apuntadorA *int
	apuntadorA = &a
	//apuntadorA := &a

	// & Dirección en memoria de la variable
	// * valor de la variable a la que se apunta
	// *& dirección en memoria de la variable que se apunta

	fmt.Println(&a)
	fmt.Println(apuntadorA)
	fmt.Println(&apuntadorA)
	fmt.Println(*apuntadorA)
	fmt.Println(*&apuntadorA)

	*apuntadorA = 200
	fmt.Println(a)

}
