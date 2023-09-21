package main

import "fmt"

func main() {
	// Initialize a slice of integers
	intSlice := []int{}
	// Fill the slice with integers from 0 to 10
	for i := 0; i <= 10; i++ {
		intSlice = append(intSlice, i)
	}
	fmt.Println(intSlice)

	for _, num := range intSlice {
		if num % 2 == 1 {
			fmt.Println(num, "is odd")
		} else {
			fmt.Println(num, "is even")
		}
	}
}