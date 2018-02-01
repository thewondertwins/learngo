# Embracing the Standard Library

In this section you’ll learn how to get the most out of Go’s amazing standard library. We’ll look at several standard library packages that are highly useful, but often ignored.  We’ll explore the common interfaces in the standard library and see how to use them effectively.  You’ll learn how to write code that is more efficient and requires fewer dependencies.

- Hidden Gems in Go’s Standard Library
- Common Interfaces
- Using std lib for maximum effect

# Hidden Gems in Go's Standard Library

## index/suffixarray

Package suffixarray implements substring search in logarithmic 
time using an in-memory suffix array.

- fast and simple string searching
- give it a slice of bytes
- search for some or all occurences of any string

How many times have you built this yourself?

## net/textproto

Package textproto implements generic support for text-based request/response 
protocols in the style of HTTP, NNTP, and SMTP.

- Build your own REDIS like server
- Build client packages for older TCP based protocols

Why didn't I know about this when I wrote https://github.com/bketelsen/handlersocket-go??

## sync/atomic

Package atomic provides low-level atomic memory primitives useful 
for implementing synchronization algorithms.

- add goroutine safe counters and configuration values to your types without using mutexes (yourself)
- Metrics! 
- modify mutable config values that are accessed in multiple goroutines.

## testing/quick

Package quick implements utility functions to help with black box testing.

- use *quick.Generator* and/or quick.Value to populate "fake"/"random" data in your types for testing/fuzzing
- *quick.CheckEqual* - test two functions to ensure that they both turn the same result with arbitrary values for each input


## time/Timer

The Timer type represents a single event. When the Timer expires, the current time will be sent on (channel) C, 
unless the Timer was created by AfterFunc. A Timer must be created with NewTimer or AfterFunc.

    type Timer struct {
        C <-chan Timer
    }

- Create time bounded operations 
- can be created with time.NewTimer() and time.AfterFunc()


## time/AfterFunc

AfterFunc waits for the duration to elapse and then calls f in its own goroutine. 
It returns a Timer that can be used to cancel the call using its Stop method.

    func AfterFunc(d Duration, f func()) *Timer

- After duration of "d", go do function "f".  
- Build automatic session expiration by using *Timer.Reset() on every new session request, delete session in func "f"

## x/net/trace (stretch, not std library)

Package trace implements tracing of requests and long-lived objects. 
It exports HTTP interfaces on /debug/requests and /debug/events

- Single server version of zipkin/dapper 

Demo: trace

# Common Interfaces

	You write too many lines of code in your programs until you learn these interfaces inside and out.

## io.MultiWriter

Read data from somewhere, write to two different places
Write web responses to network AND to response log for compliance

## io.Pipe

Combine an io.Reader and io.Writer together for synchronous buffer-free reads and writes

	stop using bytes.Buffer to store an encoded value (JSON?) and pipe directly to output

Demo:

	code $GOPATH/src/github.com/thewondertwins/learngo/material/embracestd/demos/pipe/main.go

## fmt.Stringer

One of the biggest improvements you can make to your productivity is to implement 
the Stringer interface on things you'll be outputing. 

## fmt.GoStringer

fmt.GoStringer is implemented by any value that has a GoString method, 
which defines the Go textual syntax format for that value. 
The GoString method is used to print values passed as an operand to a %#v format.
    
Great way to log/print the state of your types can be found [here](https://github.com/bbqgophers/qpid/blob/master/grill.go#L137).

# Exercise - Individual

In $GOPATH/src/github.com/thewondertwins/learngo/material/embracestd/exercises/pipe there's a copy of the pipe application we demonstrated a few minutes ago

## Assignment:

- Add an io.TeeReader so that the base64 encoded, gzipped output of the goroutine is displayed to the terminal
- HINT: you need someplace to capture the output of the 2nd reader.

## Desired Outcome:
Chaining io.Readers, io.Writers, io.TeeReader together to make 
powerful tools with no intermediate buffers
