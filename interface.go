package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHelloo(value HasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {

	person := Person{Name: "Helmi"}
	sayHelloo(person)
}
