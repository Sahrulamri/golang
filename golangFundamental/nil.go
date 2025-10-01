package main

import "fmt"


func newSlice(name string) []string {
	if name == "" {
		return nil
	} else {
		return []string{name}
	}
}
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	var person map[string]string = newMap("Eko")
	fmt.Println(person)
}