/*
Los Generics, introducidos en Go 1.18, son una característica
poderosa que permite a los programadores escribir funciones y
tipos de datos más flexibles y reutilizables.
Antes de los Generics,
Go requería que especificaras tipos de datos concretos en tus
estructuras, interfaces y funciones. Esto a menudo llevaba a la
repetición de código para manejar diferentes tipos de datos.

¿Qué son los Generics?
Los Generics permiten a los desarrolladores usar "tipos de tipo"
o "parámetros de tipo", lo que significa que puedes escribir una
función o definir una estructura que pueda operar con diferentes
tipos sin tener que reescribir la función o estructura para cada tipo.

Ventajas de usar Generics
Reutilización de Código: Escribe una función o estructura una vez,
úsala con múltiples tipos de datos.

Seguridad de Tipos: A diferencia de usar interfaces vacías
(interface{}), los Generics mantienen la seguridad de tipos
sin necesidad de conversiones de tipo constantes.

Claridad y Mantenibilidad: El código es más claro y fácil de mantener.
*/

package main

import "fmt"

// suma de enteros
func sumInt(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

// suma de flotantes
func sumFloat(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum
}

// suma entero sin signo
func sumUint(slice []uint) uint {
	var sum uint
	for _, v := range slice {
		sum += v
	}
	return sum
}

// suma float32
func sumFloat32(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println("Suma de enteros: ", sumInt(intSlice))
	fmt.Println("Suma de flotantes: ", sumFloat(floatSlice))

}
