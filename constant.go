package main

import "fmt"

func main() {
	//tidak dapat diubah lagi
	// const fName string = "helmi"
	// const lName = "helmi"
	// atau bisa ditulis
	const (
		fName string = "helmi"
		lName        = "helmi"
	)

	fmt.Println(fName)
	fmt.Println(lName)

}
