/*
Interface Segregation Principle (ISP)

Este principio nos dice que los clientes de nuestro código no deberían verse obligados
a depender de interfaces que no necesitan.

Imagina que estás en una tienda de electrónica y quieres comprar un teléfono.
No te obligan a comprar un paquete completo con teléfono, tableta, computadora y auriculares.
En lugar de eso, puedes elegir solo lo que necesitas, en este caso, solo el teléfono.

En programación, esto significa que deberíamos dividir nuestras interfaces en interfaces más pequeñas y específicas,
para que los clientes solo tengan que depender de lo que realmente necesitan.
*/

package I

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type SimplePrinter struct{}

func (sp SimplePrinter) Print() {
	fmt.Println("Imprimiendo...")
}

type SimpleScanner struct{}

func (ss SimpleScanner) Scan() {
	fmt.Println("Escaneando...")
}

type SimpleMultiFunctionDevice struct{}

func (smfd SimpleMultiFunctionDevice) Print() {
	fmt.Println("Imprimiendo...")
}

func (smfd SimpleMultiFunctionDevice) Scan() {
	fmt.Println("Escaneando...")
}

func main() {
	printer := SimplePrinter{}
	scanner := SimpleScanner{}
	multifunction := SimpleMultiFunctionDevice{}

	printer.Print()
	scanner.Scan()
	multifunction.Print()
	multifunction.Scan()
}
