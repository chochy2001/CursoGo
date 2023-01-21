package main

import "fmt"

//Defer
//Abrir y cerrar un archivo

func main() {

	defer hola1()
	adios1()
	hola1()
	hola1()
	adios1()

	fmt.Println()
	for i := 0; i <= 3; i++ {
		defer fmt.Print(i)
	}
	fmt.Println()
	for i := 0; i <= 3; i++ {
		fmt.Print(i)
	}
	fmt.Println()

}

func hola1() {
	fmt.Println("Hola")
}

func adios1() {
	fmt.Println("Adios")
}
