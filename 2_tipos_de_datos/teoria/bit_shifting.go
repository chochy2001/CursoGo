// * Bit Shifting
package main

import "fmt"

func main() {
	a := 5
	fmt.Printf("%T, %v,%b\n", a, a, a)
	b := a >> 1
	fmt.Printf("%T, %v,%b", b, b, b)

}
