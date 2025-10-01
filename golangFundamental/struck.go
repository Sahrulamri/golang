package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) sayHello() {
	fmt.Println("Hello", c.Name, "from", c.Address)
}

func main() {
	customer1 := Customer{"Eko", "Jl. Raya", 20}
	fmt.Println(customer1)
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	joko := Customer{
		Name: "Joko",
		 Address: "Jl. Raya",
		  Age: 30,
		}
	fmt.Println(joko)

	joko.sayHello()
}