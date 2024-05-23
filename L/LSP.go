/*
Liskov Substitution Principle (LSP)

Este principio nos dice que si tenemos un tipo específico, deberíamos poder usar cualquier subtipo de ese
tipo sin cambiar la forma en que funciona nuestro programa.

Por ejemplo, si tenemos una función que espera un tipo de animal,
deberíamos poder pasarle cualquier tipo de animal, ya sea un perro, un gato o un pájaro,
sin que nuestro programa falle.

Esto nos ayuda a escribir código más flexible y reutilizable.
*/

package L

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow!"
}

func LetAnimalMakeSound(animal Animal) {
	fmt.Println(animal.MakeSound())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	LetAnimalMakeSound(dog)
	LetAnimalMakeSound(cat)
}
