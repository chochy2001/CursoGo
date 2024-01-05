package main

import "fmt"

type Number interface {
	int64 | float64 | int | float32
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {

	intMap := map[string]int{"a": 1, "b": 2, "c": 3}
	floatMap := map[string]float32{"x": 1.1, "y": 2.2, "z": 3.3}

	fmt.Println("Suma de intMap: ", SumNumbers(intMap))
	fmt.Println("Suma de floatMap: ", SumNumbers(floatMap))

}
