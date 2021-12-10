package main

import "fmt"

func main() {
	slice:= []int{1,2,3,4,5,6,7,8,9,10}
	slice2:= []int{1,2,3,4,5,6,7,8,9,10}
	//salida 1 -> [1 2 3 4 5 6 7 8 9 10]
	//salida 2 -> [1 2 3 7 8 9 10]
	//salida 3 -> [2 3 7 8 9]

	fmt.Println(slice)// salida 1
	fmt.Println(append(slice[:3],slice[6:]...))// salida 2
	fmt.Println(append(slice2[1:3],slice2[6:9]...))//salida 3









}
