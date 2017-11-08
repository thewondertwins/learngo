package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	// Pass the "address of" the variable count.
	increment(&count)

	println("After: ", count, &count)
}

// increment declares inc as a pointer variable whose value is
// always an address and points to values of type int.
func increment(inc *int) {
	fmt.Println("What is inc?", inc, "a pointer")
	fmt.Println("what value is inc?", *inc, "an integer value stored in the pointer.")
	// Increment the value that the "pointer points to". (de-referencing)
	*inc++
	println("Inc:   ", *inc, &inc, inc)

}
