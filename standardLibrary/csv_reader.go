package main

import (
	"encoding/csv"
	"io"
	"strings"
	"fmt"
)

func main() {
	csvString := "eko,kurniawan,khanedy\n"+
	"budi,luhut,khanedy\n"+
	"joko,kurniawan,khanedy\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for{
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}