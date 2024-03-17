package main

import "fmt"

func main() {
	names := [...]string{"helmi", "azhar", "rendra", "gibran", "aqsa"}

	slice1 := names[2:5]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	//cara lain penulisan
	var slice3 []string = names[2:]
	fmt.Println(slice3)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "Helmi"
	newSlice[1] = "Azhar"
	//newSlice[2] = "Muhammad" //harus mengunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Asbul")
	fmt.Println(newSlice2)

	fromSlice := days[:]
	var toSlice []string = make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	thisIsArray := [...]int{1, 2, 3, 4, 5, 6}
	thisIsSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)

}
