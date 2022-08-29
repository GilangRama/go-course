package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing" || name == "Anjing"
	}

	registerUser("Gilang", blacklist)
	registerUser("Anjing", blacklist)

	registerUser("root", func(name string)bool{
		return name == "root"
	})
	registerUser("jagher", func(name string)bool{
		return name == "root"
	})
}