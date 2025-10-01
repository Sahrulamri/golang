package main

import "fmt"

func main(){
	person := map[string]string{
		"name": "Eko Kurniawan",
		"address": "Jl. Raya",
	}	
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko Kurniawan"
	book["ups"] = "Ups"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}