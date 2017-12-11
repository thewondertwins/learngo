# Pointers and Reference Types

Arrays have a fixed length and fixed data type.

- Array Example

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go
    ```

# Iterating

Iterate over the array using a for loop

- Iterate with For

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go
    ```

For is the only loop construct in Go. Use it for FOR, WHILE, DO WHILE, DO UNTIL, etc.

# Using RANGE

You can also iterate over any collection in Go using the range statement. Range is a built-in iterating function that returns the index and value of many different collection types, including arrays:

- Iterate with Range

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go
    ```

Range is very powerful, you'll use it often.

# Exercise 

Modify this example to iterate over the array using a range statement, and print the value of the array multiplied by 2.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/exercises/range

hint: do the multiplication inside the fmt.Println() function.

# Slices


An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

Unless you know that your list will contain a finite and fixed set of elements, you'll almost always use a slice when dealing with data.

- Slice Examples 

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go
    ```

# Why slices?

- Every change to an array allocates a new array. Inefficient!
- You can change a slice without an allocation
- You can operate on subsections of Arrays and Slices easily

How can you tell the difference between a slice and an array?

Slices don't have a length in the declaration:

	var sl []string // slice
	var ar [5]string // array

# Adding Data to Slices

We've shown using append to add values to a slice but you can also declare the entire slice with its values at once:

- Slice Inline Declaration

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go
    ```

# Strings

Strings are Slices of Bytes

A string is just a slice of bytes. Go has built-in support for UTF-8, and strong tools for working with non-ascii characters. While ASCII characters all take up only a single byte UTF-8 characters (or Runes) may be up to 4 bytes. Go allows you to easily handle multi-byte runes.

- Rune Example 

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/runes
	go run main.go
    ```

# Maps

A map is an _unordered_ set of values indexed by a unique key.

Maps must be initialized before they can be used.

Initialize a map with the make() function.

- Map Example 

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/maps
	go run main.go
    ```

# Maps

Map keys must define the "==" and "!=" operators. therefore you can't use functions, maps or slices as the keys in your maps.

- Maps can be declared inline too:

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/mapsinline
	go run main.go
    ```

# Map Concurrency

Maps are NOT threadsafe. It is not ok to read from a map and modify a map in two different goroutines. We'll discuss this later in the concurrency lectures. Make a mental note now. When you have a map that's being used in more than one place you need to do some sort of synchronized access.

Prior to Go 1.7, concurrent access to a map would just cause your program to be unreliable because of race conditions. In Go 1.7 and later, concurrent access to a map will cause your program to halt.

 -- TODO -- sync.Map

# Pointers

Go lets you pass function parameters by value or by reference. You will generally pass by value when the type is short lived and won't be used after the function call.

- Pass by Value Example

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyval
	go run main.go
    ```

Notice that we didn't modify the value that was passed in, we created a new integer and returned that as the result.

If you want to operate on a value and have it modified during the operation, pass it by reference using a pointer:

- Pass by Reference Example

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyref
	go run main.go
    ```

# Pointer Dereferencing

Pointer operations are similar to C family languages, but you may not do pointer arithmetic.

Use & to take the address of a variable.

Use * to de-reference a variable.







