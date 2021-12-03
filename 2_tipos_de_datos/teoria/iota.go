// * IOTA


// * Effective Go

// En Go, las constantes enumeradas se crean utilizando el enumerador iota.
//Como iota puede formar parte de una expresión y las expresiones pueden repetirse implícitamente,
//es fácil construir conjuntos de valores intrincados.

// * The Go Programming Language Specification

//Dentro de una declaración de constante, el identificador predeclarado iota representa sucesivas constantes enteras
//no tipadas. Su valor es el índice del ConstSpec respectivo en esa declaración de constante, comenzando en cero.
//Puede utilizarse para construir un conjunto de constantes relacionadas:

package main

import "fmt"

const (
	a = iota
	b
	c
)
const(
	d = iota
	e
	f
)

func main (){
	fmt.Println(a,b,c)
	fmt.Println(d,e,f)
	fmt.Printf("%T\n",a)
}
