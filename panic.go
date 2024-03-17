package main

import "fmt"

func endApp() {
	fmt.Println("end App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runnApp(err bool) {
	defer endApp()

	if err {
		panic("Ups err")
	}
}

func main() {
	runnApp(true)
	fmt.Println("Masih lanjut")
}
