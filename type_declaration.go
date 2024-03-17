package main

import "fmt"

func main() {

	type noKTP string

	var ktpHelmi = "3213123"

	var contoh string = "3242"
	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpHelmi)
	fmt.Println(contohKtp)

}
