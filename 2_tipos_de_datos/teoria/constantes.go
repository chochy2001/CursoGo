// *  Effective Go
//Las constantes en Go son sólo eso: constantes.
//Se crean en tiempo de compilación, incluso cuando se definen como locales en funciones,
//y sólo pueden ser números, caracteres (runas), cadenas o booleanos.
//Debido a la restricción en tiempo de compilación,
//las expresiones que las definen deben ser expresiones constantes, evaluables por el compilador.
//Por ejemplo, 1<<3 es una expresión constante, mientras que math.Sin(math.Pi/4)
//no lo es porque la llamada a la función math.Sin debe realizarse en tiempo de ejecución.

//En Go, las constantes enumeradas se crean utilizando el enumerador iota.
//Como iota puede ser parte de una expresión y las expresiones pueden repetirse implícitamente,
//es fácil construir conjuntos de valores intrincados.

// * The Go Programming Language Specification
// Hay constantes booleanas, constantes rúnicas, constantes enteras, constantes de punto flotante,
//constantes complejas y constantes de cadena. Las constantes rúnicas, enteras,
//de punto flotante y complejas se denominan colectivamente constantes numéricas.

//Un valor constante está representado por una runa, un entero, un punto flotante,
//un imaginario o un literal de cadena, un identificador que denota una constante,
//una expresión constante, una conversión con un resultado que es una constante,
//o el valor del resultado de algunas funciones incorporadas como unsafe.Sizeof aplicado a cualquier valor,
//cap o len aplicados a algunas expresiones, real e imag aplicados a una constante compleja
//y complex aplicados a constantes numéricas. Los valores de verdad booleanos están representados
//por las constantes predeclaradas true y false. El identificador predeclarado iota denota una constante entera.
//En general, las constantes complejas son una forma de expresión constante y se discuten en esa sección.
//Las constantes numéricas representan valores exactos de precisión arbitraria y no se desbordan.
//
//En consecuencia, no hay constantes que denoten los valores de cero negativo, infinito y no-número de IEEE-754.
//Las constantes pueden ser tipadas o no tipadas. Las constantes literales, true, false, iota, y
//ciertas expresiones constantes que contienen sólo operandos constantes no tipados son no tipadas.
//A una constante se le puede dar un tipo explícitamente mediante una declaración de constante o
//una conversión, o implícitamente cuando se utiliza en una declaración de variable o una asignación
//o como operando en una expresión. Es un error si el valor de la constante no puede representarse
//como un valor del tipo correspondiente.
//Una constante no tipada tiene un tipo por defecto que es el tipo al que se convierte implícitamente
//la constante en contextos en los que se requiere un valor tipado, por ejemplo,
//en una declaración de variable corta como i := 0 donde no hay un tipo explícito.
//El tipo por defecto de una constante no tipada es bool, rune, int, float64, complex128 o string respectivamente,
//dependiendo de si es una constante booleana, rune, entera, de punto flotante, compleja o de cadena.
//Restricción de implementación: Aunque las constantes numéricas tienen precisión arbitraria en el lenguaje,
//un compilador puede implementarlas utilizando una representación interna con precisión limitada. Dicho esto, toda implementación debe

//Representar las constantes enteras con al menos 256 bits.
//Representar las constantes de punto flotante, incluyendo las partes de una constante compleja,
//con una mantisa de al menos 256 bits y un exponente binario con signo de al menos 16 bits.
//Dar un error si no se puede representar con precisión una constante entera.
//Dar un error si no se puede representar una constante de coma flotante o compleja por desbordamiento.
//Redondear a la constante representable más cercana si no se puede representar una constante de coma flotante o
//compleja debido a los límites de precisión.
//Estos requisitos se aplican tanto a las constantes literales como al resultado de la evaluación de expresiones
//constantes.

package main

import "fmt"

const (
	nombre   string  = "Jorge"
	edad     uint8   = 20
	pi       float32 = 3.1415
	apellido string  = "Chochy"
)
var(
	anombre   string  = "Jorge"
	aedad     uint8   = 20
	api       float32 = 3.1415
	aapellido string  = "Chochy"
)

func main() {


	fmt.Println(nombre)
	fmt.Println(edad)
	fmt.Println(pi)
	fmt.Println(apellido)
	fmt.Printf("%T\n", nombre)
	fmt.Printf("%T\n", edad)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%T\n", apellido)

	anombre = "Jose"


	fmt.Println(anombre,aedad,api,aapellido)

}
