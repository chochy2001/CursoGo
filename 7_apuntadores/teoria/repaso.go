/*
Method Set de un Tipo de Valor:

	Incluye todos los métodos con un receptor de valor
	(o un receptor que es un valor).

Method Set de un Tipo de Puntero:

	Incluye todos los métodos, tanto los que tienen un receptor de valor
	como los que tienen un receptor de puntero.
*/
package main

import (
	"fmt"
)

type MyType struct{}

// Método con receptor de valor
func (m MyType) ValueMethod() {
	fmt.Println("ValueMethod")
}

// Método con receptor de puntero
func (m *MyType) PointerMethod() {
	fmt.Println("PointerMethod")
}

func main() {
	mt := MyType{}
	mp := &mt

	// Ambos métodos disponibles para el puntero
	mp.ValueMethod()
	//Lo que sucede es esto
	(*mp).ValueMethod()

	mp.PointerMethod()

	// Solo ValueMethod está disponible para el valor
	mt.ValueMethod()
	mt.PointerMethod() // Esto dará un error de compilación

	(&mt).PointerMethod() // Esto es válido
}
