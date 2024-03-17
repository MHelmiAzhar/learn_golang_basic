package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{}
	person["name"] = "helmi"
	person["address"] = "jakarta"

	//cara lain penulisan
	angka := map[int]string{
		1: "satu",
		2: "dua",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(angka)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Helmi"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
