package main

import (
	"fmt"

	
)

func endApp(){
	fmt.Println("End App")
	message := recover()
	fmt.Println(message)
}

func runApp(error bool)  {
	defer endApp()
	if error {
		panic("ERROR")
	}

	fmt.Println("Run App")
}

func main() {
	runApp(false)
	fmt.Println("End Program")
}