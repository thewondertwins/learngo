# Syntax and Types

Go has built-in types that aren't surprising.

	uint8       the set of all unsigned  8-bit integers (0 to 255)
	uint16      the set of all unsigned 16-bit integers (0 to 65535)
	uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	int8        the set of all signed  8-bit integers (-128 to 127)
	int16       the set of all signed 16-bit integers (-32768 to 32767)
	int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	float32     the set of all IEEE-754 32-bit floating-point numbers
	float64     the set of all IEEE-754 64-bit floating-point numbers
	complex64   the set of all complex numbers with float32 real and imaginary parts
	complex128  the set of all complex numbers with float64 real and imaginary parts
	byte        alias for uint8
	rune        alias for int32

# Implementation Specific Types

Implementation specific types.  (64 bit size on 64 bit platform, 32 bit on 32 bit platform)

	uint     either 32 or 64 bits
	int      same size as uint
	uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

# Non Number Types

string and bool

	string  the set of string values
	bool    a boolean (true/false) value

# Declaring and Assigning Values to Variables

Without initialization

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/withoutinit
	go run main.go


With initialization, explicit type

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/explicit
	go run main.go

With initialization, implicit type

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/implicit
	go run main.go

All three produce an integer that is indistinguishable from the others. With the implicit declaration, the compiler determines the type of the variable at compile time (not run time).

# Zero Values in Go

All builtin types have a zero value. Any allocated variable is usable even if it never has a value assigned.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/zero
	go run main.go

# Constants

Constants are variables that can't be modified at run time.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantstring
	go run main.go

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantnumber
	go run main.go

Modifying Constants Example - Fail!

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/modifyconst
	go run main.go

# Iota

Sometimes you want to declare constants that follow a sequence:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/sequence
	go run main.go

That's sort of ugly. Go gives us a compile time helper called iota that lets you skip the repetition:

Sequence with Iota

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iota
	go run main.go

Notice the difference between the two versions? Iota always starts at 0.

Skip the first value of iota

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iotaskip
	go run main.go

# Structs

A struct is a collection of fields.

Structs are types with zero or more fields.

Struct Example:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/structs
	go run main.go

Fields in a struct are called members. Reference them using a period and the field name.


# Code Organization

Go code is organized in packages. A package represents all the files in a single directory on disk. One directory can contain only files from the same package. 

You've seen this already several times. Our examples so far have all used package "main" declared at the top of the file.

# Code Organization

Source files in a package must declare the package name at the top of the file as the first code statement:

	// Package declaration
	package main

Executable programs must have a "main" package that declares a main() function:

	func main() { ...  }

Library code must declare a package name that matches the folder name it lives in. Code in folder "server" has to declare "package server".

# Scope

All variables and types declared inside a package are visible to everything else in the same package.

That means no "public" "private" "protected" modifiers.

External visibility is controlled by capitalization. Types and Functions that start with a capital letter are available outside the current package. Types and functions that start with a lower case letter are unexported, not available outside the current package.

We call this concept Exporting. A symbol that is visible outside its package is "exported".

# Package Resolution

When you installed Go earlier, you set a GOPATH environment variable in your shell.

A GOPATH is a workspace for one or more Go projects.

GOPATH is the root of the workspace and it contains three directories:

.image syntaxtypes/images/gopath.png  400 600 

# Packages

Your source code, and the code your applications depend on lives in "src".

When you build an application, it's placed in "bin".

When you compile any library, it's placed in "pkg", under a subdirectory for your computer's architecture. like pkg/darwin_amd64.

All of this is important because your GOPATH is what determines how the go compiler resolves your references to packages in code.

# Package Resolution

If your code lives at $GOPATH/src/blue/red, it's package name is "red" and you would import that code with the following statement:

	import "blue/red"

We call "blue/red" the import path of the package.

Packages that live in source code repositories like github and bitbucket have the full location of the repository as part of their import path. A project in my github repository called "captainhook" has the import path:

	"github.com/bketelsen/captainhook"

# Package Resolution

Therefore in order to use that package in your code, that package MUST live at:

	$GOPATH/src/github.com/bketelsen/captainhook

If captainhook were a library instead of an executable, when it's compiled, the compiled version of the package would be placed at:

	$GOPATH/pkg/darwin_amd64/github.com/bketelsen/captainhook
	(Assuming you compiled it on a Mac)

# Packages and GOPATH

The vast majority of developers will use one GOPATH, set it as an environment variable and forget about it. 

However, it's possible to have "clean rooms" for different sets of projects or even an individual project by simply creating a new GOPATH and setting the environment variable to that new location.

# Exercise

Exercise

Read the first half of the article here then do the exercises "Your First Program" and "Your First Library"

.link https://golang.org/doc/code.html Getting Started with Go

This article tells you to set GOPATH to $HOME/work... ignore that. *DO NOT CHANGE THE GOPATH YOU HAVE ALREADY EXPORTED*



