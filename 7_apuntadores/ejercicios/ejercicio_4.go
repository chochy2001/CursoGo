/*
Objetivo: Integrar el uso de apuntadores, method sets,
e interfaces en Go para una aplicación de gestión de empleados.

Contexto: Trabajas en el desarrollo de un sistema de gestión de empleados
que debe ser flexible para manejar diferentes tipos de empleados con distintas
responsabilidades y formas de promoción.

Tarea:

Definir Estructuras, Interfaces y Métodos:

Crea una estructura Employee con campos Name, Role y Salary.
Define una interfaz Promotable con un método Promote.
Implementa el método Promote en la estructura Employee, que actualiza el rol del
empleado y aumenta su salario.
Implementa una función Display para mostrar la información del empleado.
Función de Promoción General:

Escribe una función GeneralPromotion que acepte un Promotable y lo promueva,
demostrando el uso de interfaces.
Uso Completo de Apuntadores:

En el método main, crea una instancia de Employee, usa apuntadores para modificar
sus detalles y mostrar los cambios.
*/

package main

import "fmt"

// Employee es una estructura que representa a un empleado
type Employee struct {
	Name   string
	Role   string
	Salary float64
}

// Promotable define la capacidad de ser promovido
type Promotable interface {
	Promote(newRole string)
}

// Promote actualiza el rol del empleado y aumenta su salario
func (e *Employee) Promote(newRole string) {
	e.Role = newRole
	e.Salary *= 1.10 // 10% de aumento
}

// Display muestra la información del empleado
func (e *Employee) Display() {
	fmt.Printf("Name: %s\nRole: %s\nSalary: %.2f\n\n", e.Name, e.Role, e.Salary)
}

// GeneralPromotion promueve a cualquier empleado que sea Promotable
func GeneralPromotion(p Promotable, newRole string) {
	p.Promote(newRole)
}

func main() {
	// Crea un empleado
	employee := Employee{
		Name:   "Juan Perez",
		Role:   "Mid Golang Developer",
		Salary: 1500.00,
	}

	// Muestra la información del empleado
	employee.Display()

	// Promueve al empleado utilizando la interfaz Promotable
	GeneralPromotion(&employee, "Senior Golang Developer")

	// Muestra la información Actualizada del empleado
	employee.Display()

}
