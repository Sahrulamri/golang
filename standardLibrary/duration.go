package main

import (
	"time"
	"fmt"
)

func main() {

	var duration1 time.Duration = 10 * time.Second
	var duration2 time.Duration = 20 * time.Millisecond
	var duration3 time.Duration = 30 * time.Nanosecond
	var duration4 time.Duration = 40 * time.Minute
	var duration5 time.Duration = duration3 - duration1

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)
	fmt.Println(duration4)
	fmt.Println(duration5)
}