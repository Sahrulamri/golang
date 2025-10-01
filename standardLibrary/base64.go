package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Eko Kurniawan"
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(value)))

	fmt.Println(base64.StdEncoding.DecodeString("RWtvIEt1cnJpbmF3YW4="))
}