package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApp() {
	defer logging()
	fmt.Println("Run App")
}

func main() {
	runApp()
}
