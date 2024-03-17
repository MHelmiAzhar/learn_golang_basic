package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Helmi"
	names[2] = "Azhar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		20, 2,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])

	var Values = [...]int{
		30, 2, 456, 1,
	}

	fmt.Println(Values)

}
