// * Crea una variable de tipo string, de los dos y asignale los siguientes valores
// * en una asigna el valor de "Hola" y en la otra el valor de:
//("saludos
//a todos,
//Espero que se la esten pasando
//muy bien con el curso de Go
//")

// * Imprime los valores de las variables, con su tipo de dato
// ? Explica cuales son las diferencias entre el primer y segundo texto
// ? Como es que se interpretan los strings en la computadora

package main

import "fmt"

func main(){
	var1 := "Hola"
	var2 := `("saludos
a todos,
Espero que se la esten pasando
muy bien con el curso de go
")`

	fmt.Printf("%v,%T\n", var1,var1)
	fmt.Printf("%v,%T\n", var2,var2)
	fmt.Printf("%c %c %c %c\n",72,111,108,97)
	fmt.Printf("%d\n",65)
	fmt.Printf("%b\n",65)

}

