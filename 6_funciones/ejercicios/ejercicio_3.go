package main

import "fmt"

//Ejemplifica el uso de la palabra defer
//haciendo que una funcion se ejecute después de la función que la contiene.

func main() {
	ejemploDefer()
	fmt.Println("Dentro de Función main")
}

func ejemploDefer() {
	defer func() {
		fmt.Println("Defer dentro de la función")
	}()
	fmt.Println("Fuera de la función")

}
