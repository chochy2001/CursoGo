// * Del ejercicio anterior. Crea una variable global sin el operador de asignación corta, con el identificador de
// * EdadPersona con un tipo de dato que provenga de Edad.
// * Usa la conversión para pasar el valor guardado en edad con el tipo de dato que proviene.
// * Utiliza el operador de declaración corta para asignarle el valor a EdadPersona.
// ! Copia y pega el código del ejercicio 4

package main

import "fmt"

type Edad uint8
var edad Edad
var edadPersona uint8

func main() {
	fmt.Printf("%v\n",edad)
	fmt.Printf("%T\n",edad)
	edad = 20
	fmt.Printf("%v\n",edad)

	fmt.Printf("%v, %T\n",edadPersona,edadPersona)
	edadPersona = uint8(edad)
	fmt.Printf("%v, %T\n",edadPersona,edadPersona)
	edadPersona = 10
	fmt.Printf("%v, %T\n",edadPersona,edadPersona)


}

