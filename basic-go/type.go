package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpJag NoKTP = "11111111111"
	var marriedStatus Married = false
	fmt.Println(ktpJag)
	fmt.Println(NoKTP("22222222222"))
	fmt.Println(marriedStatus)
}