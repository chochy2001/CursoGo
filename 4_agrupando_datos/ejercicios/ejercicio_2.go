package main

import "fmt"

func main() {
	//var slice []int
	slice := make([]int, 5, 10)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scanf("%d", &slice[i])
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Println("\n")

	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}
	fmt.Printf("%T\n", slice)

}
