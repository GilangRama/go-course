package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Gilang",
		"address": "Bandung",
	}

	person["title"] = "Music Producer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// make map
	book := make(map[string]string)
	book["title"] = "Letter From Stoic"
	book["author"] = "Lupa"

	fmt.Println(book)

	book["test"] = "Salah"

	// delete(book, "test")
	fmt.Println(book)
}