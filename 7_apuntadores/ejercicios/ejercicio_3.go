/*
Objetivo: Implementar interfaces y method sets en Go.

Contexto: Eres un desarrollador de software que trabaja
en un sistema de gestión para una biblioteca. Se te ha
pedido crear un sistema que pueda manejar tanto libros impresos
como audiolibros. Ambos tipos de libros comparten algunas características,
pero también tienen sus propias particularidades.

Tarea: Define dos estructuras, PrintedBook y AudioBook.
Ambas deben implementar la interfaz Book, que tiene los siguientes métodos:

Title() string: Devuelve el título del libro.
Author() string: Devuelve el autor del libro.
DisplayInfo(): Imprime la información del libro.
Las estructuras PrintedBook y AudioBook deben
tener los siguientes campos:

PrintedBook: title, author, pageCount.
AudioBook: title, author, lengthInMinutes.
Implementa los métodos requeridos por la interfaz Book para
ambas estructuras. Para DisplayInfo, asegúrate de incluir toda la
información relevante del libro (incluyendo el número de páginas o
la duración, según corresponda).

*/

package main

import "fmt"

// Book es la interfaz que deben implementar los libros
type Book interface {
	Title() string
	Author() string
	DisplayInfo()
}

// PrintedBook es un libro impreso
type PrintedBook struct {
	title     string
	author    string
	pageCount int
}

// AudioBook es un libro de audio
type AudioBook struct {
	title           string
	author          string
	lengthInMinutes int
}

// Title devuelve el título del PrintedBook
func (p PrintedBook) Title() string {
	return p.title
}

// Author devuelve el autor del PrintedBook
func (p PrintedBook) Author() string {
	return p.author
}

// DisplayInfo imprime la información del PrintedBook
func (p PrintedBook) DisplayInfo() {
	fmt.Println("Title:", p.Title())
	fmt.Println("Author:", p.Author())
	fmt.Println("Page Count:", p.pageCount)
}

// Title devuelve el título del AudioBook
func (a AudioBook) Title() string {
	return a.title
}

// Author devuelve el autor del AudioBook
func (a AudioBook) Author() string {
	return a.author
}

// DisplayInfo imprime la información del AudioBook
func (a AudioBook) DisplayInfo() {
	fmt.Println("Title:", a.Title())
	fmt.Println("Author:", a.Author())
	fmt.Println("Length in minutes:", a.lengthInMinutes)
}

func main() {
	printed := PrintedBook{
		author:    "Hermanos Grimm",
		title:     "Blanca Nieves",
		pageCount: 24,
	}
	audio := AudioBook{
		author:          "Hermanos Grimm",
		title:           "Blanca Nieves",
		lengthInMinutes: 10,
	}
	var myBooks [2]Book
	myBooks[0] = printed
	myBooks[1] = audio

	for _, book := range myBooks {
		book.DisplayInfo()
		fmt.Println()
	}

}
