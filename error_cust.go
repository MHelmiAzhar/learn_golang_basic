package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}
	return nil
}

func main() {
	err := saveData("", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println(validationErr.Error())
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("sukses")
	}
}
