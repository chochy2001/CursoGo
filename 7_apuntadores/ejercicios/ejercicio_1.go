/*
Enunciado
Este ejercicio está diseñado para ayudar a los estudiantes a comprender
los conceptos básicos de los apuntadores en Go, incluyendo cómo asignar
valores a las variables y cómo modificar esos valores usando apuntadores.

Tarea
Asignación de Valor a una Variable:

En la función main, declara una variable numero de tipo int y
asigna un valor inicial a esta variable, por ejemplo, 5.

Crea un apuntador a la variable numero. Puedes nombrar este apuntador como ptrNumero.
Modificación del Valor Usando el Apuntador:

Utiliza el apuntador ptrNumero para cambiar el valor de numero.
Por ejemplo, puedes cambiar el valor a 10.
Imprimir los Resultados:

Imprime el valor de numero antes y después de la modificación para
demostrar que el valor ha cambiado a través del apuntador.

Ejemplo de Salida Esperada
Valor original de numero: 5
Valor de numero después de modificar a través del apuntador: 10

Objetivo de Aprendizaje
Entender cómo los apuntadores permiten acceder y modificar el
valor de las variables a las que apuntan.

Apreciar la diferencia entre cambiar el valor directamente y
usar un apuntador para cambiarlo.
*/

package main

import "fmt"

func main() {
	//Asignamos un valor a la variable numero
	number := 5
	fmt.Println("Valor original de numero:", number)

	//Creamos un apuntador a la variable numero
	ptrNumber := &number

	//Modificamos el valor de la variable numero a través del apuntador
	*ptrNumber = 10
	fmt.Println("Valor de numero después de modificar a través del apuntador:", number)

}
