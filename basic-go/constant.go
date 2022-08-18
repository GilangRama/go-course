package main

import "fmt"

func main() {
	const (
		firstName = "Gilang" // tidak bisa diubah karena konstan
		lastName  = "Mahardhika"
		age       = 22
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}