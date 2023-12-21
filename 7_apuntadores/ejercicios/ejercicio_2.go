/*
Enunciado
Este ejercicio tiene como objetivo ayudar a los estudiantes a
comprender el uso y la validez de varios operadores relacionados
con apuntadores en Go. Los operadores a explorar son *, &, &*, **, y &&.
Algunos de estos operadores son válidos en Go, mientras que otros no lo son.
Los estudiantes deberán determinar cuáles son válidos, cuáles no, y por qué.
*/

//Operador * (Dereferenciación):
//Utiliza el operador * para acceder al valor al que apunta un apuntador.

//Operador & (Dirección):
//Utiliza el operador & para obtener la dirección de memoria de una variable.

//Combinación &*:
//Explora qué sucede cuando se combinan los operadores & y *.
//¿Es válido en Go? ¿Qué significa si es válido?

//Operador ** (Doble Dereferenciación):
//Intenta crear un apuntador a un apuntador (doble apuntador)
//y utiliza ** para acceder al valor original.
//Investiga si Go permite este tipo de operación.

//Operador && (Doble Dirección):
//Evalúa la validez de obtener la dirección de la dirección de una variable (&&).
//¿Es posible o tiene sentido en Go?

package main

import "fmt"

func main() {
	value := 10

	//uso del operador & para obtener la dirección de memoria de una variable
	ptr := &value
	fmt.Println("Dirección de memoria de value:", ptr)

	//uso del operador * para acceder al valor al que apunta un apuntador
	valueDeferenced := *ptr
	fmt.Println("Valor obtenido a traves de *ptr:", valueDeferenced)

	//Combina los operadores & y * para obtener el valor original de value
	redundantDirection := &*ptr
	fmt.Println("Valor obtenido a traves de &*ptr:", redundantDirection)

	//Creando apuntador doble
	ptrToPtr := &ptr
	triplePtr := &ptrToPtr
	doubleDereferenced := **ptrToPtr
	fmt.Println("Valor obtenido a traves de **ptrToPtr:", doubleDereferenced)
	fmt.Println("Valor obtenido a traves de ***triplePtr:", ***triplePtr)

	//Usando && - No es valido
	//doubleDirection := &&value //No es valido
	//fmt.Println("Valor obtenido a traves de &&value:", doubleDirection)

	cadenaA := "Hola"
	cadenaB := "Mundo"
	ptrA := &cadenaA
	ptrB := &cadenaB

	intercambiarCadenas(ptrA, ptrB)
	fmt.Println(*ptrA, *ptrB) // Imprime: Mundo Hola
}

func intercambiarCadenas(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}
