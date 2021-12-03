package main

import "fmt"

func main() {
	fmt.Println("Operadores logicos")
	// && -> AND -> Y
	fmt.Println("\n&& -> AND -> Y")
	fmt.Println("true && true =", true && true)     //true
	fmt.Println("true && false =", true && false)   //false
	fmt.Println("false && true =", false && true)   //false
	fmt.Println("false && false =", false && false) //false

	// || -> OR -> O
	fmt.Println("\n|| -> OR -> O")
	fmt.Println("true || true =", true || true)     //true
	fmt.Println("true || false =", true || false)   //true
	fmt.Println("false || true =", false || true)   //true
	fmt.Println("false || false =", false || false) //false

	// !
	fmt.Println("\n! -> NOT -> NEGACION")
	fmt.Println("!true =", !true)   //false
	fmt.Println("!false =", !false) //true

}
