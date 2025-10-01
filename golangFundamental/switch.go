package main

import "fmt"

func main() {
	var nilai = 80

	switch nilai {
	case 80:
		fmt.Println("Nilai A")
	case 70:
		fmt.Println("Nilai B")
	case 60:
		fmt.Println("Nilai C")
	case 50:
		fmt.Println("Nilai D")
	default:
		fmt.Println("Nilai E")
	}

	switch length := len("Eko"); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}