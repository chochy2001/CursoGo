package main

import "fmt"

func main() {
	// * Crea un arreglo que contenga 5 números enteros.
	// ! (pedidos por el usuario)
	// * Imprime los valores del arreglo con su indice
	// for init, condicion, incremento {
	// for range

	arreglo := [5]int{}
	for i := 0; i < len(arreglo); i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scanf("%d", &arreglo[i])
	}

	for i, v := range arreglo {
		fmt.Println(i, v)
	}
	fmt.Println("\n")

	for i := 0; i < len(arreglo); i++ {
		fmt.Println(i, arreglo[i])
	}
	fmt.Printf("%T\n", arreglo)

}
