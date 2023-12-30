package main

import (
	"fmt"
)

type Sumable2 interface {
	~int | ~float64 | ~uint | float32 | uint8 | ~string
urs

func Sum2[T Sumable2](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	//Es un alias para un tipo de dato int
	type Entero int
	//Es un alias para un tipo de dato float64
	type Flotante float64
	//Es un alias para un tipo de dato uint
	type EnteroSinSigno uint
	//Es un alias para un tipo de dato string
	type Cadena string

	//slice de Entero
	intSlice := []Entero{1, 2, 3, 4, 5}
	//slice de Flotante
	floatSlice := []Flotante{1.1, 2.2, 3.3, 4.4, 5.5}
	//slide de EnteroSinSigno
	uintSlice := []EnteroSinSigno{1, 2, 3, 4, 5}
	//slice de Cadena
	stringSlice := []Cadena{"hola", "mundo"}

	//print Sum Entero
	fmt.Println("Suma de enteros: ", Sum2(intSlice))
	fmt.Println("Suma de flotantes: ", Sum2(floatSlice))
	fmt.Println("Suma de enteros sin signo: ", Sum2(uintSlice))
	fmt.Println("Suma de strings: ", Sum2(stringSlice))

}
