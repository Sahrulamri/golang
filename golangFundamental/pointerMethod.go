package main

import "fmt"

type Man struct {
	Name string
	
}

func (man *Man) Maried() {
	man.Name = "Mr . " + man.Name
}

func main() {
	eko := Man{"Eko Kurniawan"}
	eko.Maried()
	fmt.Println(eko.Name)
}