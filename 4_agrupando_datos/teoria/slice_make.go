package main

import "fmt"

func main() {
	// De vuelta a la asignación. La función incorporada make(T, args) tiene un propósito diferente al de new(T).
	//Sólo crea cortes, mapas y canales, y devuelve un valor inicializado (no puesto a cero) de tipo T (no *T).

	//La razón de la distinción es que estos tres tipos representan, bajo cuerda,
	//referencias a estructuras de datos que deben ser inicializadas antes de su uso.
	//Un slice, por ejemplo, es un descriptor de tres elementos que contiene un puntero a los datos
	//(dentro de un array), la longitud y la capacidad, y hasta que esos elementos se inicialicen, el slice es nulo.
	//Para las rebanadas, los mapas y los canales, make inicializa la estructura de datos interna y prepara el valor
	//para su uso. Por ejemplo,

	//make([]int, 10, 100)

	//asigna un array de 100 ints y luego crea una estructura slice con una longitud de 10
	//y una capacidad de 100 que apunta a los primeros 10 elementos del array.
	//(Cuando se hace una rebanada, la capacidad puede ser omitida; ver la sección sobre rebanadas para más información).
	//Por el contrario, new([]int) devuelve un puntero a una estructura de rebanada recién asignada y puesta a cero,
	//es decir, un puntero a un valor de rebanada nulo.
	slice := make([]int, 10, 12)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice[0] = 1
	slice[2] = 12
	fmt.Println(slice)
	slice[9] = 13
	fmt.Println(slice)

	slice = append(slice, 14)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 14)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 14)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
