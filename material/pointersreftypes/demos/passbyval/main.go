// Sample program to show the basic concept of using a pointer
// to share data.
package main

// main is the entry point for the application.
func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	// Pass the "address of" the variable count.
	count2 := increment(count)

	println("After: ", count, &count)
	println("After: ", count2, &count2)
}

// increment adds one to the value passed in and
// returns the result
func increment(inc int) int {
	inc++
	return inc
}
