// * Crea tu propio tipo de dato que provenga de un entero para la edad (Lo más optimo posible)
// * Crea una variable con tu nuevo tipo con el identificador "edad" sin el operador de declaración corta
// * Imprime el valor de la variable "edad", el tipo y asignale el valor de 20
// * imprime el valor de edad.

package main

import "fmt"

/*type Edad uint8
var edad Edad
 */


func main() {
	fmt.Printf("%v\n",edad)
	fmt.Printf("%T\n",edad)
	edad = 20
	fmt.Printf("%v\n",edad)
}
