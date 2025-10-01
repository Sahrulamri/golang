package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Eko Kurniawan"}
	sayHello(&person)

	animal := Animal{"Kucing"}
	sayHello(animal)
}