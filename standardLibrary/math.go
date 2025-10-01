package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.67))
	fmt.Println(math.Floor(1.67))
	fmt.Println(math.Round(1.67))
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
}