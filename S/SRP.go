/*
Single Responsibility Principle (SRP)

Este principio nos dice que cada cosa en nuestro código debería hacer solo una cosa.

Imagina que estás construyendo una casa. En lugar de tener un solo trabajador que haga todos los trabajos,
como construir paredes, instalar fontanería y hacer la electricidad, sería mejor tener diferentes trabajadores,
cada uno especializado en una tarea específica.

En programación, esto significa que cada función, estructura o tipo debería tener una sola razón para cambiar,
una sola tarea que realiza; Esto hace que nuestro código sea más fácil de entender y mantener.
*/

package S

import "fmt"

// Printer tiene la responsabilidad de imprimir mensajes.
type Printer struct{}

func (p Printer) PrintMessage(message string) {
	fmt.Println(message)
}

// Logger tiene la responsabilidad de registrar mensajes en un archivo de registro.
type Logger struct{}

func (l Logger) LogMessage(message string) {
	// Implementación para escribir en un archivo de registro
	fmt.Println("Registrando mensaje:", message)
}

func main() {
	printer := Printer{}
	printer.PrintMessage("¡Hola Mundo!")

	logger := Logger{}
	logger.LogMessage("Este es un mensaje de registro")
}
