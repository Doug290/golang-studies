package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range ints {
		if ((number%2) != 0) {
			fmt.Printf("The number %v is Odd", number)
			fmt.Println()
		} else {
			fmt.Printf("The number %v is Even", number)
			fmt.Println()
		}
	}
}
