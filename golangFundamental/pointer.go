package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = new(Address)
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	
}