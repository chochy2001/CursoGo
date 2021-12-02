package main

import "fmt"

func main() {
	// Imprime todos los numeros del 0 al 10 cuyo residuo es igual a 0 dividido entre 2
	/*
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}

		// Imprime todos los fúmeros del 0 al 100 multiplos de 3, pero no los múltiplos de 5.
		for i := 0; i <= 100; i++ {
			if i%3 == 0 && i%5 != 0 {
				fmt.Println(i)
			}
		}
		// Imprime todos los números del 1 al 100 multiplos de 2, pero no los múltiplos de 3.
		for i := 1; i <= 100; i++ {
			if i%2 == 0 && i%3 != 0 {
				fmt.Println(i)
			}
		}
		// Imprime el residuo de los numeros del 1 al 100 divididos entre 3 y 5.
		for i := 1; i <= 100; i++ {
			//fmt.Println(i % 3, i % 5)
			fmt.Println(i)
			fmt.Printf("residuo de 3: %d, residuo de 5: %d\n",i % 3, i % 5)
		}
	*/

	// Imprime el residuo de los numeros del 1 al 100 divididos entre 2 y 3.
	for i := 1; i <= 100; i++ {
		//fmt.Println(i % 3, i % 5)
		fmt.Println(i)
		fmt.Printf("residuo de 2: %d, residuo de 3: %d\n", i%2, i%3)
	}
}
