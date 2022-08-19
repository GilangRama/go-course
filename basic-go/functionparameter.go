package main

import "fmt"

func sayHello(fname string, lname string, age int) {
	fmt.Println("Hello", fname, lname,"Age", age)
}

func main() {
	sayHello("Gilang", "Mahardhika", 22)
}