package main

import "fmt"

func main() {

	name := "helmi az"

	switch name {
	case "helmi":
		fmt.Println("helmi")
	case "azhar":
		fmt.Println("azhar")
	case "muhammad":
		fmt.Println("muhammad")

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
