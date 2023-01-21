package main

import "fmt"

func main() {
	sliceEnteros := []int{1, 2, 3, 4}
	total := sumaNumeros123(sliceEnteros...)
	fmt.Println("Total:", total)
	totalPares := numerosPares(sumaNumeros123, sliceEnteros...)
	fmt.Println("Total Pares:", totalPares)

}

func sumaNumeros123(numeros ...int) int {
	sumaTotal := 0
	fmt.Printf("%T\n", numeros)
	for _, n := range numeros {
		sumaTotal += n
	}
	return sumaTotal
}

func numerosPares(sum func(numeros ...int) int, numeros ...int) int {
	var sliceEnteros []int
	for _, n := range numeros {
		if n%2 == 0 {
			sliceEnteros = append(sliceEnteros, n)
		}
	}
	return sum(sliceEnteros...)
}
