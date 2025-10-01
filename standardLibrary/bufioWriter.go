package main

import (
	"bufio"
	// "fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Eko Kurniawan\n")
	_, _ = writer.WriteString("Joko Kurniawan\n")
	_ = writer.Flush()
}