// Sample program to show how to takes slices of slices to create different
// views of and make changes to the underlying array.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	var sl []string
	fmt.Println("Empty Slice")
	fmt.Println("*************************")
	inspectSlice(sl)
	sl = append(sl, "Mango")
	fmt.Println("Slice after Append")
	fmt.Println("*************************")
	inspectSlice(sl)

	// Create a slice with a length of 5 elements and a capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	fmt.Println("Fruit Slice")
	fmt.Println("*************************")
	inspectSlice(slice1)

	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are [starting_index : (starting_index + length)
	slice2 := slice1[2:4]

	fmt.Println("Slice of Fruit Slice")
	fmt.Println("*************************")
	inspectSlice(slice2)

	// Change the value of the index 0 of slice3.
	slice2[0] = "CHANGED"

	// Display the change across all existing slices.
	fmt.Println("Fruit Slice")
	fmt.Println("*************************")
	inspectSlice(slice1)

	fmt.Println("Slice of Fruit Slice")
	fmt.Println("*************************")
	inspectSlice(slice2)

	slice2 = append(slice2, "Kiwi")
	slice2 = append(slice2, "Lychee")
	inspectSlice(slice2)
	inspectSlice(slice1)

	// Another slice of slice1. Showing index shortcuts
	// [index:] takes everything from index to end
	// Don't forget that slices and arrays are 0 indexed!
	// This takes items 3,4,5 from the original slice
	slice3 := slice1[2:]
	fmt.Println("Shortcut Slice")
	fmt.Println("*************************")
	inspectSlice(slice3)

	// Another slice of slice1. Showing index shortcuts
	// [:index] takes everything up to index
	// [:3] == give us the last three elements
	slice4 := slice1[:3]
	fmt.Println("Shortcut Slice 2")
	fmt.Println("*************************")
	inspectSlice(slice4)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, v := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			v)
	}
	fmt.Println("\n")
}
