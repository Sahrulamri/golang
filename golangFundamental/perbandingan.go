package main

import "fmt"

func main() {
	var name1 = "Eko Kurniawan"
	var name2 = "Eko Kurniawan"

	var result = name1 == name2
	fmt.Println(result)

	var name3 = "eko kurniawan"
	var name4 = "Eko Kurniawan"
	var result2 = name3 == name4
	fmt.Println(result2)

	var abjad1 = "A"
	var abjad2 = "a"
	var result3 = abjad1 > abjad2
	fmt.Println(result3)
}