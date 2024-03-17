package main

import "fmt"

func main() {
	//dia mengakses variable di luar functionnya

	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
}
