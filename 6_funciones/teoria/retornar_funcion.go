package main

import "fmt"

func main() {
	fmt.Printf("%T\n", retornarFuncion)
	fmt.Printf("%T\n", retornarFuncion())
	fmt.Printf("%T\n", retornarFuncion()())

}

func hola123() string {
	s := "hola"
	return s
}

func retornarFuncion() func() int {
	return func() int {
		return 1
	}
}
