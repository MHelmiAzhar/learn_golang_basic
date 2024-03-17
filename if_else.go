package main

import "fmt"

func main() {
	name := "azhar"

	if name == "helmi" {
		fmt.Println("helmi")
	} else if name == "azhar" {
		fmt.Println("azhar")
	} else {
		fmt.Println("none")
	}

	//if else short statment
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
