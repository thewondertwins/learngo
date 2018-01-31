# Write Code Like the Go Team

In this chapter of the book you’ll learn how to make your code look like it belongs in Go’s standard library, and why that’s important.  

Topics:

- how to organize your code into packages, and what those packages should contain  
- code patterns and conventions that are prevalent in the standard library  
- how to write your code to be more clear and understandable  
- unwritten Go conventions that go beyond “go fmt” and make you look like a veteran Go contributor wrote it


# Section One Outline

In this section we'll cover three main areas that will help you write code that looks more professional, performs better, and is easier to read and understand.

- Packages
- Naming Conventions
- Source Code Conventions

# Packages

## Code Organization

The organization of your code is an area of frequent confusion even among seasoned Go developers.  

There are two key areas of code organization in Go that will make a huge impact on the usability, testability, and functionality of your code:

- Package Naming
- Package Organization

We can't talk about them separately because one influences the other.

## Package Organization - Libraries

Let's start with Go as our canonical reference. Here's a list of the top level packages in Go 1.7's standard library:

	archive   cmd       crypto    errors    go        index     math      path      sort      syscall   unicode
	bufio     compress  database  expvar    hash      internal  mime      reflect   strconv   testing   unsafe
	builtin   container debug     flag      html      io        net       regexp    strings   text      vendor
	bytes     context   encoding  fmt       image     log       os        runtime   sync      time

Immediately it is clear:

	Packages contain code that has a single purpose

When a group of packages provides a common set of functionalities with different implementations, they're organized under a parent.  Look at the contents of package *encoding*:

	ascii85     base32      binary      encoding.go hex         pem
	asn1        base64      csv         gob         json        xml

The *encoding* package has only one file *encoding.go* which describes a common interface that all the implementation packages will implement.  

Using this naming structure, it's very clear that *encoding/json* is a package that manages encoding for the *json* format.

Some commonalities:

- Packages names describe their purpose
- It's very easy to see what a package does by looking at the name
- Names are generally short
- When necessary, use a descriptive parent package and several children implementing the functionality -- like the *encoding* package

## Package Organization - Applications

The packages we've seen are all libraries. They're intended to be used by imported and used by some executable program like a service or command line tool.

What should the organization of your executable applications look like?

When you have an application, the package organization is subtly different.  The difference is the command, the executable that ties all of those packages together.

Application package organization has a huge impact on the testability and functionality of your system.

	When writing an application your should aim to write code that is easy to understand, easy to refactor, and easy for someone else to maintain.

Most libraries focus on providing a singularly scoped function; logging, encoding, network access.

Your application will tie all of those libraries together to create a tool or service. That tool or service will be much larger in scope.  

When you're building an application, you should organize your code into packages, but those packages should be centered on two categories:

- Domain Types
- Services

*Domain Types* are types that model your business functionality and objects. 

*Services* are packages that operate on or with the domain types.

[Standard package layout by Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)

A domain type is the substance of your application. If you have an inventory application, your domain types might include *Product* and *Supplier*. If you have an HR administration system, your domain types might include *Employee*, *Department*, and *Business*Unit*.

The package containing your domain types should also define the interfaces between your domain types and the rest of the world. These interfaces define the things you want to do with your domain types.

- *ProductService*  
- *SupplierService*
- *AuthenticationService*
- *EmployeeStorage*
- *RoleStorage*

Your domain type package should be the root of your application repository.  This makes it clear to anyone opening the codebase what types are being used, and what operations will be performed on those types.

The domain type package, or *root* package of your application should not have any external dependencies. It exists for the sole purpose of describing your types and their behaviors.

	The implementations of your domain interfaces should be in separate packages, organized by dependency.

Dependencies include:

- External data sources
- Transport logic (http, RPC)

You should have one package per dependency.

Why one package per dependency?

- Simple testing
- Easy substitution/replacement
- No circular dependencies

We'll dive deeper into the implications of package organization in the *Interfaces* and *Testing* chapters.


# Conventions

## Naming Conventions

	there are two hard things in computer science: 
	cache invalidation, naming things, and off-by-one errors
*Source*: Every developer on Twitter


Naming things *is* hard, but putting some thought into your type, function, and package names will make your code more readable.


## Naming Conventions - Packages

A package name should have the following characteristics:

- short

	Prefer "transport" over "transportmechanisms"

- clear 

	Name for clarity based on function: "bytes" 
	Name to describe implementation of external dependency: "postgres" 

Packages should provide functionality for one and only one purpose.  Avoid *catchall* packages:

	util
	helpers
	etc.

Frequently they're a sign that you're missing an interface somewhere.

	util.ConvertOtherToThing() should probably be refactored into a Thinger interface

*catchall* packages are always the first place you'll run into problems with testing and circular dependencies, too.

## Naming Conventions - Variables

Some common conventions for variable names:

- use camelCase not snake_case
- use single letter variables to represent indexes
	- for i:=0; i < 10; i++ {}
- use short but descriptive variable names for other things
	- var count int
	- var cust Customer

There are no bonus points in Go for obfuscating your code by using unnecessarily short variables. Use the scope of the variable as your guide. 

	The farther away from declaration you use it, the longer the name should be.

Here are examples of using the variable's scope as a guide to the length of the variable name:

.link https://github.com/golang/go/blob/master/src/fmt/format.go#L12
.link https://github.com/golang/go/blob/master/src/fmt/format.go#L326
.link https://github.com/golang/go/blob/master/src/flag/flag.go#L293

- use repeated letters to represent a collection/slice/array
	- var tt []*Thing

- inside a loop/range, use the single letter
	- for i, t := range tt {}

These conventions are common inside Go's own source code.

## Naming Conventions - Functions and Methods

Avoid a package-level function name that repeats the package name.  

	GOOD:  log.Info()
	BAD:   log.LogInfo()

The package name already declares the purpose of the package, so there's no need to repeat it.

Go code doesn't have setters and getters.

	GOOD:  custSvc.Customer()
	BAD:   custSvc.GetCustomer()

## Naming Conventions - Interfaces

If your interface has only one function, append "-er" to the function name:

	type Stringer interface{
		String() string
	}

If your interface has more than one function, use a name to represent its functionality:

	type CustomerStorage interface {
		Customer(id int) (*Customer, error)
		Save(c *Customer)  error
		Delete(id int) error
	}

Some purists think that all interfaces should end in `-er`.  I think interfaces should be descriptive and readable.  

## Naming Conventions - Source Code

Inside a package separate code into logical concerns.

If the package deals with multiple types, keep the logic for each type in its own source file:

	src/goteam/exercises/inventory/postgres

	orders.go
	suppliers.go
	products.go

In the package that defines your domain objects, define the types and interfaces for each object in the same source file:

	src/goteam/exercises/inventory

	orders.go -- contains Orders type and OrderStorage interface


# Smaller Tips

- Make comments in full sentences, always.

	```
	// An Order represents an order from a customer.	
	type Order struct {}
	```
- Use `goimports` to manage your imports, and they'll always be in canonical order. Standard lib first, external next. 

- Avoid the `else` clause.  Especially in error handling.  
	```
	if err != nil {
		// error handling
		return // or continue, etc.
	} 
	// normal code
	```

# Coding Like the Go Team

Following these conventions will make your source code easier to read, easier to maintain, and easier for someone else to understand.


# Exercise - Individual

In *$GOPATH/src/github.com/thewondertwins/learngo/material/goteam/exercises/inventory/* there's a skeleton application with all the interfaces designed.

*Assignment:*

Create an appropriate package structure under the inventory package for the applications with the following assumptions:

- Database used for storage of domain objects will be postgres
- Supplier providing the SupplierService will be "Acme"
- Two methods of service transport:  REST over HTTP and RPC over TCP

You only need to create the directories representing your packages.  You do not need to implement any code.

*Desired Outcome:*

Understand the implications of dependencies when creating packages
