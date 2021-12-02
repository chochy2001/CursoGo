package main

import "fmt"

func main() {

	x:= []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[4:])
	fmt.Println(x[:3])
	fmt.Println(x[:5])
	fmt.Println(x[2:8])
	fmt.Println(x[:9])

}
