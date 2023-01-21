package main

import "fmt"

var xyz int

func main() {
	fmt.Println(xyz)
	hola1234()
	{
		var xyzz = 9
		fmt.Println(xyzz)
		xyzz = 10
		fmt.Println(xyzz)
	}
	adios1234()

}

func hola1234() {
	xyz++
	fmt.Println(xyz)
}

func adios1234() {
	xyz++
	fmt.Println(xyz)
}
