package main

import "fmt"

func main() {
	var name string

	name = "Sahrul Amri"
	fmt.Println("Hello", name)

	name = "Eko Kurniawan"
	fmt.Println("Selamat Sore", name)

	var benda = "Kursi"
	fmt.Println("Benda ini adalah", benda)

	hewan := "Kucing"
	fmt.Println("Hewan ini adalah", hewan)

	var(
		firstName = "Eko"
		lastName = "Kurniawan"
	)

	fmt.Println(firstName, lastName)
}