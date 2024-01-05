package main

// any es un alias para interface {} 1.18
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Swap2[T interface{}](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	a, b := 1, 2
	println(a, b)
	Swap(&a, &b)
	println(a, b)

	c, d := "hola", "mundo"
	println(c, d)
	Swap(&c, &d)
	println(c, d)

}
