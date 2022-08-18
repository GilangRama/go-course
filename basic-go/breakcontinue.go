package main

import "fmt"

func main() {

	// Break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Test", i)
	}

	// continue

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Continue", i)
	}

	// fizzbuzz
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}