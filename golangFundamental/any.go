package main

import "fmt"

func ups() any {
	return 1
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}