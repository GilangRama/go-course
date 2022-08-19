package main

import "fmt"

func getFullName() (string, string, int) {
	return "Gilang", "Mahardhika", 22
}

func main() {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)
}