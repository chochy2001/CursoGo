package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//salida 1 -> [2 3 4 5]
	//salida 2 -> [1 2 3 4 5 6 7 8 9 10]
	//salida 3 -> [6 7 8 9 10]
	//salida 4 -> [7 8 9]
	//salida 5 -> [1]
	fmt.Println(slice[1:5])
	//fmt.Println(slice[0:10])
	fmt.Println(slice[:])
	//fmt.Println(slice[5:10])
	fmt.Println(slice[5:])
	fmt.Println(slice[6:9])
	fmt.Println(slice[:1])
	//fmt.Println(slice[0:1])

}
