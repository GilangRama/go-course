package main

import "fmt"

func main() {
	name := "Rama"

	if name == "Gilang" {
		fmt.Println("Betul")
	} else if name == "Rama" {
		fmt.Println("Masih Betul")
	} else {
		fmt.Println("Salah")
	}

	// short statement
	address := "Bandung Sumedang"
	// length := len(address)
	if length := len(address); length > 5 {
		fmt.Println("Panjang teuing")
	} else {
		fmt.Println("Aman")
	}
}