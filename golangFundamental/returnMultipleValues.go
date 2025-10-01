package main

import "fmt"

func getFullname() (string, string) {
	return "Eko", "Kurniawan"
}

func main() {
	firstName, lastName := getFullname()
	fmt.Println(firstName, lastName)

	firstName2, _ := getFullname()
	fmt.Println(firstName2)
}