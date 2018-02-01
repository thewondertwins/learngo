# Real World Testing

In this section we’ll use real-world code from open source Go projects as we learn how to write code that is testable.  We’ll show how to create packages, functions and types that are more testable, and show you how to test those hard to test corners of your code that are often ignored.  You’ll learn advanced techniques that will give you more confidence when you ship your code.

- Writing Testable Code
- Interfaces and Mocks
- Integration Tests
- Testing Hard to Test Code
- Use Table Driven Tests


# General Testing Advice

- One test file per source file
- Use internal tests where possible

# Internal vs. External Tests

Internal tests share the same package name as the package being tests.  External tests use the package name with `_test` at the end.

External tests see the tested package as if it were an external package, so they follow the same rules of visibility that Go's packages do.   

In general, use internal tests unless an external test is required to break a dependency cycle.

In Go's standard lib, this only happens 3 times.  It should be a sign that your package layout is too complex if you require this.


## External Tests - with Go Examples

Contrary to the above advice, if you're writing Example tests:

[pkg/testing](https://golang.org/pkg/testing/)

They look best in an external test package so you can demonstrate the package identifiers.

[examples](https://blog.golang.org/examples)

# Writing Testable Code

There are several keys to writing testable code:

- Dependency Injection
- Interfaces
- Keep functions small

# Dependency Injection

This is where it all starts to tie together.  The first key to writing testable code is to pass your dependencies into the New() "constructors" of your types:

[kubernetes/servicecontroller](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/service/servicecontroller.go#L100)


	Accept (or store) an interface, return a concrete type

This makes testing simpler because you can test your types with many different scenarios

[kubernetes/servicecontroller_test](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/service/servicecontroller_test.go#L36)

# Interfaces

Where practical, don't accept a file, or a string as a function parameter. Accept an interface.

[uber/encoder](https://github.com/uber-go/zap/blob/master/encoder.go#L36)

This Uber logging library has Encoders which write different style log messages.  Instead of accepting a file to write the logs in the WriteEntry() function, it accepts an io.Writer.

[uber/encoder_test](https://github.com/uber-go/zap/blob/master/json_encoder_test.go#L55)

Making it possible to test logging with a bytes.Buffer instead of writing to a file and reading the results


# Keep Functions Small

Functions should be small and do one thing.  This makes it easier to test the function without testing side-effects of other pieces of your code:

[errors/stack](https://github.com/pkg/errors/blob/master/stack.go)

[errors/stack_test](https://github.com/pkg/errors/blob/master/stack_test.go)

# Keep Test Self-Contained

Go can and will run more than one test at a time.  All of your tests should be self-contained.

- Use small functions to set up required state for each test
- Use small functions to tear down any state
- Use temp files auto-generated instead of fixed file names for test data
- Use :0 for your port on any listeners to get a randomly assigned TCP port


## Examples:

[syncthing/auditservice_test](https://github.com/syncthing/syncthing/blob/master/cmd/syncthing/auditservice_test.go#L18)

## Keep Test Self-Contained - TestMain 

The opposite of self-contained tests is TestMain

[TestMain](https://golang.org/pkg/testing/)

Setup and Teardown is generally overkill but there are scenarios when you need to test something that requires heavy setup, or exact conditions that must be controlled perfectly or run in a specific sequence.

Use TestMain for this.

[hackergo/missing_test](https://github.com/machenbach/hackergo/blob/master/missing_test.go#L8)

Here it's used to simulate a file being piped through standard input. e.g. cat myfile > thisapp

[ldbl/start_test](https://github.com/safronizator/ldbl/blob/master/start_test.go#L8)

# Interfaces and Mocks

Use interfaces and mocks together to allow you to test functionality without crossing package boundaries.

[noms/chunk_store.go](https://github.com/attic-labs/noms/blob/master/go/chunks/chunk_store.go#L15)

[noms/caching_chunk_haver_test.go](https://github.com/attic-labs/noms/blob/master/go/datas/caching_chunk_haver_test.go#L16)

This entire project is an excellent example of interfaces and mocks for testing

[github.com/attic-labs/noms](https://github.com/attic-labs/noms)


# Integration Tests

Integration tests assume you're testing your code against some or all external services in a testing/staging environment.

[Testing and validation](http://peter.bourgon.org/go-in-production/#testing-and-validation)

Use build flags

	go test -tags=integration

coupled with environment variables or command line flags to point to your staging/testing environment's dependency servers (database, elastic search, etc.)


# Testing Hard to Test Code

- Command line flag parsing: TestMain
- Find the right interface, it's probably already there:
	
	example: testing reading a config file - take an io.Reader 
	Production code opens a file and sends that to config.Parse()
	Test code uses a byte buffer

## Table Driven Tests

Table driven tests make your test files cleaner, easier to read and easier to maintain.

inputs, got, want

[errors/errors_test.go](https://github.com/pkg/errors/blob/master/errors_test.go#L11)

[kubernetes/bearertoken_test.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/auth/authenticator/bearertoken/bearertoken_test.go#L67)

[backlog/level_test.go](https://github.com/backplaneio/backlog/blob/master/level_test.go#L29) 

[ora/rsetCfg_test.go](https://github.com/rana/ora/blob/master/rsetCfg_test.go#L10)

# Wrapup

If you can write it, you can test it, but it may not always be easy.  For inspiration look at test suites of large applications in Go like Kubernetes, Docker, noms.

# Exercise

	go get github.com/attic-labs/noms
	cd $GOPATH/src/github.com/attic-labs/noms
	
Explore the code first.  Very complete test suite, uses Testify for testing.

Choose some packages and run `go test` in them.  Read the test source code and see how they use mocks and interfaces together to enable package-level testing independently.