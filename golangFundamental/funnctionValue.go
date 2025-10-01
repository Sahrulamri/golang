package main

import "fmt"

func getGoodbay(name string) string {
	return "Goodbay " + name
}

func main() {
	// name := "Eko Kurniawan"
	goodbay := getGoodbay
	fmt.Println(goodbay("Eko Kurniawan"))
}