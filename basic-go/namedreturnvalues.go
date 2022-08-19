package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Gilang"
	middleName = "Rama"
	lastName = "Mahardhika"

	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}