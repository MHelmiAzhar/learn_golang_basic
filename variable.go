package main

import "fmt"

func main() {
	// var name string
	// name = "helmi azhar"

	//atau bisa seperti ini tulisannya
	name := "helmi azhar" // := hanya untuk inisialisasi pertama

	fmt.Println(name)

	name = "Muhammad Helmi"
	fmt.Println(name)

	var (
		fName = "helmi"
		lName = "azhar"
	)

	fmt.Println(fName)
	fmt.Println(lName)
}
