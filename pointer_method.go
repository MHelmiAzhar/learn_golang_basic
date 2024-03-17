package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	helmi := Man{"helmi"}
	helmi.Married()

	fmt.Println(helmi.Name)
}
