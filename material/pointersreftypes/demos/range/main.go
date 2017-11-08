// Sample program to show how to declare and iterate over
// arrays of different types.
package main

import "fmt"

// main is the entry point for the application.
func main() {

	// Declare an array of 4 integers that is initialized
	// with some values.
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers.
	for i, num := range numbers {
		fmt.Println(i, num)
	}

	// Iterate over the array of numbers.
	// Ignore the index, we don't need it
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Iterate over the array of numbers.
	// Ignore the value, we don't need it, just get the index
	for idx := range numbers {
		fmt.Println(idx)
	}
}
