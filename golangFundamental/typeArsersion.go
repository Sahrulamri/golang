package main

import "fmt"

func random() any {
	return "Eko"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Type is string", value)
	case int:
		fmt.Println("Type is int")
	default:
		fmt.Println("Type is unknown")
	}


}