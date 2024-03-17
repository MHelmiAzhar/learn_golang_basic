package main

import "fmt"

type Custumer struct {
	Name    string
	Address string
	Age     int
}

// method
func (customer Custumer) sayHello(name string) {
	fmt.Println("hello", name, "my name is ", customer.Name)
}

func main() {
	var helmi Custumer
	helmi.Name = "Muhammad Helmi Azhar"
	helmi.Address = "Jakarta"
	helmi.Age = 20

	fmt.Println(helmi)

	rahman := Custumer{
		Name:    "Rahman",
		Address: "Bogor",
		Age:     20,
	}
	fmt.Println(rahman)
	budi := Custumer{"Budi", "Jakarta", 20}
	fmt.Println(budi)

	budi.sayHello("Agus")
}
