package main

import "fmt"

func main() {
	name := "gher"

	switch name {
	case "Jagher":
		fmt.Println("Betul")
	case "Gilang":
		fmt.Println("Betul Juga")
	default:
		fmt.Println("Salah")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Aman")
	case false:
		fmt.Println("Bukan duh")
	}

	// tanpa kondisi
	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Terlalu Panjang")
	case length > 5:
		fmt.Println("Sudah Aman")
	default:
		fmt.Println("Kependekan")

	}

}