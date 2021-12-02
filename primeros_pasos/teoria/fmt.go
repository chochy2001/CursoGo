package main

func main() {
	//var edad = 20
	//var booleano = true
	// * Funciona para cualquier valor
	//! %v	the value in a default format
	//! when printing structs, the plus flag (%+v) adds field names

	//! %#v	a Go-syntax representation of the value
	//! %T	a Go-syntax representation of the type of the value
	//! %%	a literal percent sign; consumes no value

/*
	fmt.Println("Esto es para cualquier tipo de dato")
	fmt.Printf("(tipo) T =  %T\n", edad)
	fmt.Printf("(valor) v = %v\n", edad)
	fmt.Printf("(valor GO) #v = %#v\n",edad)
 */

	// * Funciona para enteros
	//! %b	base 2
	//! %c	the character represented by the corresponding Unicode code point
	//! %d	base 10
	//! %o	base 8
	//! %O	base 8 with 0o prefix
	//! %q	a single-quoted character literal safely escaped with Go syntax.
	//! %x	base 16, with lower-case letters for a-f
	//! %X	base 16, with upper-case letters for A-F
	//! %U	Unicode format: U+1234; same as "U+%04X"

	/*
	fmt.Println("\nEsto es para los enteros")
	fmt.Printf("(base dos (binario) b = %b\n", edad)
	fmt.Printf("(caracter) c = %c\n",edad)
	fmt.Printf("(base diez (decimal)) d = %d\n",edad)
	fmt.Printf("(base ocho (octal)) o = %o\n",edad)
	fmt.Printf("(base ocho (octal)) O = %O\n",edad)
	fmt.Printf("(Comillas) q = %q\n",edad)
	fmt.Printf("(base 16) x = %x\n",edad)
	fmt.Printf("(base 16) X = %X\n",edad)
	fmt.Printf("(Unicode) U = %U\n",edad)
	 */


	// * Solo funciona para booleanos
	//! %t	the word true or false
	/*
	fmt.Println("\nEsto es para los booleanos")
	fmt.Printf("(booleano) t = %t\n",booleano)
	 */

	// * Funciona para flotantes y complejos
	//var flotante = 12.5
	//var complejo = complex(12.5,12.5)
	//! %b	decimalless scientific notation with exponent a power of two,
	//!	in the manner of strconv.FormatFloat with the 'b' format,
	//!	e.g. -123456p-78
	//! %e	scientific notation, e.g. -1.234456e+78
	//! %E	scientific notation, e.g. -1.234456E+78
	//! %f	decimal point but no exponent, e.g. 123.456
	//! %F	synonym for %f
	//! %g	%e for large exponents, %f otherwise. Precision is discussed below.
	//! %G	%E for large exponents, %F otherwise
	//! %x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	//! %X	upper-case hexadecimal notation, e.g. -0X0.23ABCP+20
	/*
	fmt.Println("Numeros flotantes")
	fmt.Printf("(flotante)%%b =  %b\n",flotante)
	fmt.Printf("(flotante)%%e =  %e\n",flotante)
	fmt.Printf("(flotante)%%E =  %E\n",flotante)
	fmt.Printf("(flotante)%%f =  %f\n",flotante)
	fmt.Printf("(flotante)%%F =  %F\n",flotante)
	fmt.Printf("(flotante)%%g =  %g\n",flotante)
	fmt.Printf("(flotante)%%G =  %G\n",flotante)
	fmt.Printf("(flotante)%%x =  %x\n",flotante)
	fmt.Printf("(flotante)%%X =  %X\n",flotante)

	fmt.Println("\nNumeros Complejos")
	fmt.Printf("(complejo)%%b =  %b\n",complejo)
	fmt.Printf("(complejo)%%e =  %e\n",complejo)
	fmt.Printf("(complejo)%%E =  %E\n",complejo)
	fmt.Printf("(complejo)%%f =  %f\n",complejo)
	fmt.Printf("(complejo)%%F =  %F\n",complejo)
	fmt.Printf("(complejo)%%g =  %g\n",complejo)
	fmt.Printf("(complejo)%%G =  %G\n",complejo)
	fmt.Printf("(complejo)%%x =  %x\n",complejo)
	fmt.Printf("(complejo)%%X =  %X\n",complejo)
	 */

	// * Funciona para strings y slice of bytes
	//! %s	the uninterpreted bytes of the string or slice
	//! %q	a double-quoted string safely escaped with Go syntax
	//! %x	base 16, lower-case, two characters per byte
	//! %X	base 16, upper-case, two characters per byte
	/*
	var nombre = "Jorge"
	fmt.Println("\nstring")
	fmt.Printf("%%s = %s\n",nombre)
	fmt.Printf("%%q = %q\n",nombre)
	fmt.Printf("%%x = %x\n",nombre)
	fmt.Printf("%%X = %X\n",nombre)
	 */
	// ? Diferencias entre print, printf y println
	//var imprimir = 20

	//var hola = "Hola"
	/*fmt.Println(imprimir)
	fmt.Print(imprimir)
	fmt.Printf("\n%T, %b, %v, %d, %o, %x",hola,imprimir,imprimir,imprimir,imprimir,imprimir)
	*/

	// \a   U+0007 alert or bell
	// \b   U+0008 backspace
	// \f   U+000C form feed
	// \n   U+000A line feed or newline
	// \r   U+000D carriage return
	// \t   U+0009 horizontal tab
	// \v   U+000B vertical tab
	// \\   U+005C backslash
	// \'   U+0027 single quote  (valid escape only within rune literals)
	// \"   U+0022 double quote  (valid escape only within string literals)
	// fmt.Printf("\n%T,\b%b,\f%v,\r%d,\t%o,\v%x  \"\"",hola,imprimir,imprimir,imprimir,imprimir,imprimir)
	// ? Diferencias entre sprint, sprintf y sprintln
	/*
	fmt.Print(imprimir)
	valor := fmt.Sprintf("%b",imprimir)
	fmt.Print(valor)
	 */
	// ? Diferencias entre fprint, fprintf y fprintln




}
