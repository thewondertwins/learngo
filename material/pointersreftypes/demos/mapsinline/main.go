// Sample Map program
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare and make a map uses a slice of users as the key.
	// This allocates an empty map.
	grades := map[string]int{
		"susan": 100,
		"bob":   89,
	}

	fmt.Println("Map Values")
	fmt.Println("**********")
	// Iterate over the map.
	for key, value := range grades {
		fmt.Println(key, value)
	}

}
