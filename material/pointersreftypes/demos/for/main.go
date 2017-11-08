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
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}
