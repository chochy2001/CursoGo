package main

import "fmt"

// Concrete types

type Person struct {
	Name string
	Age  int
}

var p Person

// Interface types
type Speaker interface {
	Speak() string
}

func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Robot struct{}

func (r Robot) Speak() string {
	return "Beep beep!"
}

func main() {
	var fido Dog
	var rob Robot

	Greet(fido)
	Greet(rob)

}
