package main

import (
	"regexp"
	"fmt"
)

func main() {
	var regex = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}`)
	fmt.Println(regex.MatchString("Eko Kurniawan"))

	fmt.Println(regex.FindString("Eko Kurniawan"))
	fmt.Println(regex.FindAllString("Eko Kurniawan", -1))
}