package main

import "fmt"

type Speaker interface {
	speak() string
}

type Animal struct {
	name string
}

func (a Animal) speak() string {
	return a.name + "make a sound."
}

type Dog struct {
	Animal
}

func (d Dog) speak() string {
	return d.name + "barks."
}
func saySomething(s Speaker) {
	fmt.Println(s.speak())
}
func main() {
	animal := Animal{name: "some animal"}

	dog := Dog{Animal{name: "Rex"}}

	saySomething(animal)
	saySomething(dog)
}
