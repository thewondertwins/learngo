// Sample program to show how to declare and iterate over
// arrays of different types.
package main

import "fmt"

// main is the entry point for the application.
func main() {

	// Declare an array of 4 integers that is initialized
	// with some values.
	numbers := [4]int{10, 20, 30, 40}

	// Go doesn't allow unused variables to exist in a program.
	// the compiler will complain on this exercise until you add
	// your assignment.  You may remove this.
	fmt.Println(numbers)

	// Exercise:
	// modify this example to iterate over the array using a range statement,
	// and print the value of the array element multiplied by 2.
	fmt.Println(numbers)
}
