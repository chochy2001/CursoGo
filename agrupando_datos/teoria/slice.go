package main

import "fmt"

func main() {
	// ! Slices
	// * Effective Go
	//Las rebanadas envuelven a los arrays para dar una interfaz más general,
	//potente y conveniente a las secuencias de datos.

	//A excepción de los elementos con dimensión explícita, como las matrices de transformación,
	//la mayor parte de la programación de arrays en Go se realiza con rebanadas en lugar de arrays simples.
	//Las rebanadas contienen referencias a un array subyacente,

	//y si asignas una rebanada a otra, ambas se refieren al mismo array.
	//Si una función toma un argumento en forma de porción, los cambios que se realicen
	//en los elementos de la porción serán visibles para la persona que llama,
	//de forma análoga a pasar un puntero a la matriz subyacente. Por lo tanto,
	//una función de lectura puede aceptar un argumento de trozo en lugar de un puntero y un recuento;
	//la longitud dentro del trozo establece un límite superior de la cantidad de datos a leer.



	// * Language Specification
	//Un slice es un descriptor de un segmento contiguo de un array subyacente y proporciona acceso a una
	//secuencia numerada de elementos de ese array. Un tipo de trozo denota el conjunto de todos los trozos
	//de matrices de su tipo de elemento. El número de elementos se denomina longitud del segmento y nunca es negativo.
	//El valor de una rebanada no inicializada es nulo.

	x:= []int{1,2,3,4,5,6,7,8,9,10,}
	fmt.Println(x)

	// * Composite literals
	// Los literales compuestos construyen valores para structs, arrays, slices y maps y
	//crean un nuevo valor cada vez que se evalúan.
	//Están formados por el tipo del literal seguido de una lista de elementos entre llaves.
	//Cada elemento puede ir precedido opcionalmente por una clave correspondiente.

}
