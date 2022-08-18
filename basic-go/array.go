package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Gilang"
	names[1] = "Rama"
	names[2] = "Mahardhika"

	fmt.Println(names[0], names[1], names[2])

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values[0], values[1], values[2])

	values[0] = 12

	fmt.Println(len(names), len(values))
	fmt.Println(values[0], values[1], values[2])
}