package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 5
	var c = a + b * d / e
	fmt.Println(c)

	var h = 10
	h += 10
	fmt.Println(h)
	h += 34
	fmt.Println(h)

	var i = 20
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
	
}