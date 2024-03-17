package main

import "fmt"

type Address struct {
	City, Provience, Country string
}

func changeCountryToIndoensia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}

	changeCountryToIndoensia(&address)
	fmt.Println(address)
}
