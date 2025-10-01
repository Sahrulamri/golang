package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddress(address *Address) {
	address.Country = "Bandung"
}
func main() {
	address := Address{}
	changeAddress(&address)

	fmt.Println(address)
}