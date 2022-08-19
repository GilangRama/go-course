package main

import "fmt"

func sayHello(words string) string {
	if words == "" {
		return "Hi "
	} else {
		return "Hi " + words
	}
}

func main() {
	result := sayHello("Jagher")
	fmt.Println(result)

	fmt.Println(sayHello(""))
}