package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "ID is required"}
	} else if data == nil {
		return &notFoundError{Message: "Data not found"}
	} else {
		return nil
	}
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		fmt.Println(err.Error())
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println(validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println(notFoundErr.Message)
		} else {
			fmt.Println("Unknown error", err.Error())
		}
	} else {
		fmt.Println("Data saved")
	}
}