package main

import(
	"encoding/csv"
	// "fmt"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Eko", "Kurniawan", "Khanedy"})
	_ = writer.Write([]string{"Budi", "Luhut", "Khanedy"})
	_ = writer.Write([]string{"Joko", "Kurniawan", "Khanedy"})

	writer.Flush()
}