package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80
	fmt.Println(lulusNilaiAkhir, lulusAbsensi)

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
	fmt.Println(nilaiAkhir > 80 && absensi > 80)
}