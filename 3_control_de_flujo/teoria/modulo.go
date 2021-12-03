package main

import "fmt"

// % -> residuo de una division
// 10/5 = 2
// 10%5 = 0
func main() {
	fmt.Println("Números divisibles en 2 entre 0 y 100")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Números divisibles en 3 entre 0 y 100")
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}



}
