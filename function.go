package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello")
}

func sayHelloTo(fName string, lName string) {
	fmt.Println("Hello", fName, lName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Helmi", "Azhar"
}

func getComplateName() (fName, mName, lName string) {
	fName = "Muhammad"
	mName = "Helmi"
	lName = "Azhar"

	return fName, mName, lName
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// function as value
func getGoodbye(name string) string {
	return "good bye " + name
}

// function as parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// anonymous func
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are Blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursive func
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {

	sayHello()
	sayHello()
	sayHello()

	sayHelloTo("Helmi", "Azhar")

	result := getHello("helmi")
	fmt.Println(result)

	fName, lName := getFullName()

	fmt.Println(fName, lName)

	firstName, _ := getFullName()
	fmt.Println(firstName)

	a, b, c := getComplateName()
	fmt.Println(a, b, c)

	fmt.Println(sumAll(10, 11, 12))

	numbers := []int{12, 3, 5}
	fmt.Println(sumAll(numbers...))

	contoh := getGoodbye
	misal := getGoodbye

	fmt.Println(contoh("Gibran"))
	fmt.Println(misal("Rendra"))

	sayHelloWithFilter("aqsa", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	//anonymous func

	registerUser("Rahman", func(name string) bool {
		return name == "Anjing"
	})
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})

	fmt.Println(factorialLoop(3))
	fmt.Println(factorialRecursive(3))
}
