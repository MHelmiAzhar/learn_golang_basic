package main

import (
	"fmt"
)

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke-", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke-", counter)
	// }

	names := []string{"Muhammad", "helmi", "azhar"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index =", index, name)
	}

	fmt.Println("selesai")
}
