package main

import "fmt"

func main() {
	var name1 = "Gilang"
	var name2 = "Gilang"
	var name3 = "Jagher"

	var num1 = 10
	var num2 = 10
	var num3 = 100
	var num4 = 9

	fmt.Println(name1 == name2)
	fmt.Println(name1 == name3)

	fmt.Println(num1 == num2)
	fmt.Println(num2 == num3)
	fmt.Println(num3 > num4)
	fmt.Println(num4 < num1)
}