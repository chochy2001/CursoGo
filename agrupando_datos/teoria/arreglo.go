package main

import "fmt"

func main() {
	var años [3]int
	fmt.Println(años)
	años[0] = 2001
	fmt.Println(años)

	años[2] = 2025
	fmt.Println(años)
	años[1] = 205
	fmt.Println(años)
	años[0] = 2020
	fmt.Println(años)

	hola := []string{"hola", "mundo", "!"}
	fmt.Println(hola)
	fmt.Println(len(hola))
	adios := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 120, 123, 23, 243, 45, 4, 6545, 654, 45, 634, 563, 345, 634, 5634, 56, 76, 54, 63, 5, 2}
	fmt.Println(adios)
	fmt.Println(len(adios))
	var edad [10]int
	fmt.Println(edad)
	// * Effective Go
	// Las matrices son útiles cuando se planifica la disposición detallada de la memoria
	//y a veces pueden ayudar a evitar la asignación, pero principalmente son un bloque de
	//construcción para las rebanadas, el tema de la siguiente sección.
	//Para sentar las bases de ese tema, he aquí unas palabras sobre las matrices.

	//Hay grandes diferencias entre la forma en que funcionan los arrays en Go y en C. En Go

	//Los arrays son valores. Asignar un array a otro copia todos los elementos.
	//En particular, si pasas un array a una función, ésta recibirá una copia del array, no un puntero a él.
	//El tamaño de un array es parte de su tipo. Los tipos [10]int y [20]int son distintos.
	//La propiedad del valor puede ser útil pero también costosa; si quieres un comportamiento y eficiencia como en C,
	//puedes pasar un puntero al array.

	// * Language Specification
	// Tipos de matrices
	//Una matriz es una secuencia numerada de elementos de un solo tipo, llamado tipo de elemento.
	//El número de elementos se denomina longitud de la matriz y nunca es negativo.

}
