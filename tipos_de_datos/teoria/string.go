// * Tipo de Dato String
// A string type represents the set of string values.
// A string value is a (possibly empty) sequence of bytes.
// The number of bytes is called the length of the string and is never negative.
// Strings are immutable: once created, it is impossible to change the contents of a string.
// The predeclared string type is string; it is a defined type.

// The length of a string s can be discovered using the built-in function len.
//The length is a compile-time constant if the string is a constant.
//A string's bytes can be accessed by integer indices 0 through len(s)-1.
//It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid.



//Un tipo de cadena representa el conjunto de valores de cadena.
//Un valor de cadena es una secuencia (posiblemente vacía) de bytes.
//El número de bytes se denomina longitud de la cadena y nunca es negativo.
//! Las cadenas son inmutables: una vez creadas, es imposible cambiar su contenido.
//El tipo de cadena predeclarado es string; es un tipo definido.

//La longitud de una cadena s puede descubrirse mediante la función incorporada len.
//La longitud es una constante en tiempo de compilación si la cadena es una constante.
//Se puede acceder a los bytes de una cadena mediante los índices enteros 0 a len(s)-1.
//Es ilegal tomar la dirección de un elemento de este tipo; si s[i] es el byte i's de una cadena, &s[i] no es válido.

// * Traducción realizada con la versión gratuita del traductor www.DeepL.com/Translator

package main

import "fmt"

func main(){
	text := "Hola"
	differentText := `"Hola"`
	fmt.Println(text)
	fmt.Printf("%T\n", text)
	fmt.Println(len(text))
	fmt.Println("--------------")


	fmt.Println(differentText)
	fmt.Printf("%T\n", differentText)
	fmt.Println(len(differentText))

	// * slice de bytes

	sliceBytes := []byte(differentText)
	fmt.Println(sliceBytes)
	fmt.Println(len(sliceBytes))

	sliceBytes[0] = 10
	fmt.Println(sliceBytes)
	fmt.Println(len(sliceBytes))
	fmt.Printf("%T\n", sliceBytes)



}

