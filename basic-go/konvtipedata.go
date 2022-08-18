package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai64)
	var nilai8 = int8(nilai16)

	fmt.Println(nilai32, nilai64, nilai16, nilai8)

	var name = "Gilang Rama Mahardhika"
	var e = name[0]
	var eString = string(e) 

	fmt.Println(name, e, eString)
}