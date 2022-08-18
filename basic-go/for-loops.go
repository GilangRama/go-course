package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	slice := []string{"Gilang", "Rama", "Mahardhika"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Jagher", "Gilang", "Rama", "Mahardhika"}
	for i, name := range names {
		fmt.Println("index", i, "name =", name)
	}

	// map
	person := make(map[string]string)
	
	person["name"] = "Jagher"
	person["age"] = "22"

	for key, value := range person {
		fmt.Println("key", key, "value", value)
	}
}