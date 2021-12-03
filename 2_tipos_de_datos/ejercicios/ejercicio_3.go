// * Crea constantes con un tipo de dato especifico y sin tipo de dato especifico. *
// * Explica cuales son las diferencias entre ambas y cuando usar que tipo de constantes
// * imprime el valor de las constantes, su tipo de dato
// ! y su direccion de memoria.

package main

import "fmt"

const (
	x     = 20 // Su tipo de dato es int en este momento y puede ser cualquier otro
	y int = 30 // Su tipo de dato es int en este momento y no puede ser otro
)

func main() {
	hola := 10
	fmt.Printf("tipo de dato: %T y su valor: %v\n", x, x)
	fmt.Printf("tipo de dato: %T y su valor: %v\n", y, y)
	fmt.Scanf("%d", &hola)
	fmt.Println(&hola)
}
