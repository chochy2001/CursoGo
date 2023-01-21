package main

import "fmt"

func main() {
	/*sliceInt := []int{1, 2, 3, 4, 5}
	sumaNumeros2(sliceInt...)
	sumaNumeros3(sliceInt)
	*/
	strings("Jose", "Andrea", "Juan", "Pedro")
	sliceString := []string{"Jose", "Andrea", "Juan", "Pedro"}
	strings(sliceString...)
}

func sumaNumeros2(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	suma := 0
	for _, v := range x {
		suma += v
	}
	fmt.Println(suma)

}
func sumaNumeros3(x []int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	suma := 0
	for _, v := range x {
		suma += v
	}
	fmt.Println(suma)
}

func strings(x ...string) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
