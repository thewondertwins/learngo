// Sample Map program
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare and make a map uses a slice of users as the key.
	// This allocates an empty map.
	grades := make(map[string]int)

	// Assign some values
	grades["brian"] = 100

	fmt.Println("Map")
	fmt.Println("**********")
	// Iterate over the map.
	for key, value := range grades {
		fmt.Println(key, value)
	}

	// Change a value
	grades["brian"] = 65

	// allocate a new variable with the value of the
	// map at index "brian"
	g := grades["brian"]
	fmt.Printf("brian's grade is %d\n", g)

	// Add more values
	grades["bob"] = 85
	grades["dipendra"] = 88

	fmt.Println("\nAdding Values")
	fmt.Println("**********")
	// Iterate over the map.
	for key, value := range grades {
		fmt.Println(key, value)
	}

}
