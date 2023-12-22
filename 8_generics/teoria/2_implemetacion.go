package main

import "fmt"

type Sumable interface {
	int | float64 | uint | float32 | string | uint8
}

func Sum[T Sumable](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	float32Slice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	uintSlice := []uint{1, 2, 3, 4, 5}
	stringSlice := []string{"hola", "mundo"}

	fmt.Println("Suma de enteros: ", Sum(intSlice))
	fmt.Println("Suma de flotantes: ", Sum(floatSlice))
	fmt.Println("Suma de flotantes 32: ", Sum(float32Slice))
	fmt.Println("Suma de enteros sin signo: ", Sum(uintSlice))
	fmt.Println("Suma de strings: ", Sum(stringSlice))

}
