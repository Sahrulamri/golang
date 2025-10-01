package main

import "fmt"
import "slices"

func main() {
	names := []string{"Eko", "Kurniawan", "Khanedy"}
	values := []int{1, 2, 3}

	fmt.Println(names, values)
	fmt.Println((slices.Min(names)))
	fmt.Println((slices.Max(values)))
	fmt.Println((slices.Contains(names, "Eko")))
}