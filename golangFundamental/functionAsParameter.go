package main

import "fmt"

func sayHellowithFilter(name string, filter func(string) string) string {
	nameFiltered := filter(name)
	return "Hellow " + nameFiltered
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello := sayHellowithFilter("Eko Kurniawan", func(name string) string {
		return name[0:5]
	})
	fmt.Println(sayHello)

	sayHellowithFilter("Anjing", spamFilter)
}