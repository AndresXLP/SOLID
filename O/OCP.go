/*
Open/Closed Principle (OCP)

Este principio nos dice que nuestro código debería estar abierto para ser extendido pero cerrado para ser modificado.

Piensa en una caja de herramientas. Puedes agregar nuevas herramientas a la caja en cualquier momento,
pero no necesitas cambiar la caja en sí para hacerlo.

En programación, esto significa que deberíamos poder agregar nuevas funcionalidades a nuestro código sin
tener que cambiar el código existente. Esto se logra mediante la creación de interfaces y la composición en lugar de
modificar directamente el código.
*/

package O

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(shape Shape) {
	fmt.Println("Área:", shape.Area())
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	PrintArea(rectangle)
	PrintArea(circle)
}
