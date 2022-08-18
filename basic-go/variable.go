package main

import "fmt"

func main() {
	var name string

	name = "Gilang"
	fmt.Println(name)

	name = "Jagher"
	fmt.Println(name)

	var surname = "Jagher Gilang"
	fmt.Println(surname)

	var age = 22
	fmt.Println(age)

	occupation := "Produser"
	fmt.Println(occupation)

	occupation = "Programmer"
	fmt.Println(occupation)

	var (
		firstName = "Gilang"
		lastName = "Mahardhika"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}