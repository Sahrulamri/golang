package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var NilaiAbsensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = NilaiAbsensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}