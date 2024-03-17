package main

import "fmt"

type Address struct {
	City, Provience, Country string
}

func main() {
	// address1 := Address{"Jaksel", "Jakarta", "Indonesia"}
	// address2 := address1 //copy value

	// address2.City = "Jakarta Selatan"
	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Jaksel", "Jakarta", "Indonesia"}
	address2 := &address1 //POINTER

	address2.City = "Jakarta Selatan"
	fmt.Println(address1)
	fmt.Println(address2)
}
