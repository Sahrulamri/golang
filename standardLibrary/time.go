package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc  := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-01-01 00:00:00"
	valuetime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valuetime)
	}

	fmt.Println(valuetime.Format("2006-01-02 15:04:05"))
	fmt.Println(valuetime.Add(24 * time.Hour).Format("2006-01-02 15:04:05"))

	fmt.Println(valuetime.Year())
	fmt.Println(valuetime.Month())
	fmt.Println(valuetime.Day())
	fmt.Println(valuetime.Hour())
	fmt.Println(valuetime.Minute())
	fmt.Println(valuetime.Second())
	fmt.Println(valuetime.Nanosecond())
}