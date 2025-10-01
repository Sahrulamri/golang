package main

import "fmt"

func main() {
	name := "Eko Kurniawan"

	if name == "Eko Kurniawan" {
		fmt.Println("Hello Eko")
	 } else if name == "Kurniawan" {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Hi")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}