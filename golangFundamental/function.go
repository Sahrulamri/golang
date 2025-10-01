package main 

import "fmt"

func sayhelo(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(sayhelo("Eko"))
}