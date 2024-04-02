package main

import "fmt"

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Walking"
}

func AnimalAction(a Animal) string {
	return fmt.Sprintf("says: %s\nmoves by: %s\n", a.Speak(), a.Move())
}

func main() {
	dog := Dog{}
	fmt.Println("Dog:")
	fmt.Println(AnimalAction(dog))

	fmt.Println()

	cat := Cat{}
	fmt.Println("Cat:")
	fmt.Println(AnimalAction(cat))
}
