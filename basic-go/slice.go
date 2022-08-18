package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	var slice2 = months[6:9]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	months[5] = "Test ubah"
	fmt.Println(slice1)


	// append
	days := [...]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	days1 := days[3:]
	fmt.Println(days1)

	days2 := append(days1, "Waduh libur")
	fmt.Println(days2)

	days2[1] = "Test duh ya gimana"
	fmt.Println(days2)
	fmt.Println(days1)

	// new slice
	newSlice := make([]string, 4, 8)

	newSlice[0] = "Jagher"
	newSlice[1] = "Gilang"
	newSlice[2] = "Rama"
	newSlice[3] = "Mahardhika"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}