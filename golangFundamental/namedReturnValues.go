package main

import "fmt"

func getCompleteName() (firstName string, lastName string) {
	firstName = "Eko"
	lastName = "Kurniawan"
	return firstName, lastName
}

func main() {
	firstName, lastName := getCompleteName()
	fmt.Println(firstName, lastName)
}