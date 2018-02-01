# Wrangling Concurrency

In this section we’ll explore advanced techniques to enable concurrency in your applications.  You’ll learn when and how to use channels and mutexes, and explore common pitfalls of each option.  You’ll learn how to detect and prevent data races and common concurrency errors, and you’ll get practical experience using advanced tools like sync.WaitGroup and sync.ErrorGroup to make your concurrent applications more reliable.

- Channels
- sync.WaitGroup
- sync.ErrorGroup
- Mutexes
- Race Detection


# Channels

## Channel Refresher

*channels* are like pipes that connect goroutines, allowing them to send data back and forth.

- make a new unbuffered channel:
	
	events := make(chan string)

- send to a channel:

    events <- "Brian Logged In"

- receive from a channel

    e := <- events

Sends and Receives will BLOCK until both the sender and receiver are ready.  For unbuffered channels, this means that you can't send until there is something ready to receive.

## Channel Refresher - Unbuffered Channels

	func main() {
		events := make(chan string)

		events <- "Brian Logged In"

		e := <-events
		fmt.Println(e)
	}  

Will this work? 

NO

main() executes in one goroutine, sequentially, top to bottom.  An unbuffered channel can't send until there is something ready to receive.

*Quick fix*: Send in a goroutine

	func main() {
		events := make(chan string)

		go func() {
			events <- "Brian Logged In"
		}()

		e := <-events
		fmt.Println(e)
	}

## Channel Refresher - Buffered Channels

- make a new buffered channel:

    ```
	events := make(chan string, 1) // buffers up to 1 send
	```
	```
	func main() {
		events := make(chan string, 1)

		events <- "Brian Logged In"

		e := <-events
		fmt.Println(e)
	} 
	```

Not terribly practical, but proves the point of buffers.  A buffered channel can accept up to BUFFER SIZE writes before blocking, even when there is no receiver ready.

## Channel Refresher 

*Every* problem I've ever had with concurrency in Go boils down to those two concepts.

- Unbuffered channels block until both sides are ready
- Buffered channels block after buffer is full

## Concurrency 

Go's concurrency revolves around goroutines, either with or without channels. Without channels there is no communication. With channels there can be communication of both values and state.

Let's review a few patterns for goroutines and channels.

## Concurrency - Channel as Signal

	func expensiveProcess(done chan bool) {
		// simulate long task
		time.Sleep(5 * time.Second)
		done <- true
	}

	func main() {
		done := make(chan bool)
		go expensiveProcess(done)
		<-done
	}  

main() will block on the _<-done_ receive until the goroutine sends on the channel.  This signals that the goroutine is finished.

	func main() {
		a := make(chan string)
		go func() {
			for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
				a <- strings.ToUpper(os.Args[i])
			}
			close(a)
		}()

		for s := range a {
			fmt.Println(s)
		}
	}  

You can use *range* to iterate over values in a channel until it's closed.  An unclosed channel will block, breaking the range. Close to signal you're done sending values.

# sync.WaitGroup

## Wait For Goroutine Completion

Use sync.WaitGroup to count launched goroutines and wait for them to declare that they're done.

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("I'm done")
	}()
	wg.Wait()

- Most commonly used to launch goroutines in a loop.  Perhaps ranging over a list of work items.
- Use `defer wg.Done()` at the top of your goroutine so you don't forget to announce the finish of the goroutine

## sync.WaitGroup example

	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
		i := i // WHY?  closure and goroutine!
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(strings.ToUpper(os.Args[i]))
		}()
	}
	wg.Wait()


- Nice method of accounting for launched goroutines and waiting for them to complete
- No communication without adding channels! Can't pass back values or errors.

## sync.WaitGroup example with Channel

	func main() {
		uppers := make(chan string, len(os.Args)-1)
		var wg sync.WaitGroup
		for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
			i := i // WHY?  closure and goroutine!
			wg.Add(1)
			go func() {
				defer wg.Done()
				uppers <- strings.ToUpper(os.Args[i])
			}()
		}
		wg.Wait()
		close(uppers)
		for u := range uppers {
			fmt.Println(u)
		}
	} 

- What if your goroutine can return an error?
- Make a result type that has the return value and an error.
- Yuck!

## sync.WaitGroup example with Result channel

	type result struct {
		number int
		err    error
	}

	func main() { // convert input strings into integers
		results := make(chan result, len(os.Args)-1)
		var wg sync.WaitGroup
		for i := 1; i < len(os.Args); i++ { // arg[0] is exe name
			i := i // WHY?  closure and goroutine!
			wg.Add(1)
			go func() {
				defer wg.Done()
				num, err := strconv.Atoi(os.Args[i])
				results <- result{number: num, err: err}
			}()
		}
		wg.Wait()
		close(results)
		for r := range results {
			fmt.Println(r.number, r.err)
		}
	}

# sync.ErrGroup

Finally there's a better way!

[sync.ErrGroup](https://godoc.org/golang.org/x/sync/errgroup)

- Synchronization like sync.WaitGroup
- Adds error propagation
- Adds cancelation by using `context` package

Big Example:

[bketelsen/gogrep](https://github.com/bketelsen/gogrep)

gogrep walkthrough

	go get github.com/bketelsen/gogrep

# Mutexes

A mutex is a Mut(ually)Ex(clusive) Lock

- Lighter weight and easier to reason about than channels.
- Use when sharing objects that need to be modified by many goroutines

My most common use:

- Have a type with a map in it? Add a mutex to protect the map.

Mutexes get a bad rap, but I use them a lot.  Remember the Go Proverb:

	Don't communicate by sharing memory, share memory by communicating.

If all you need to do is change a piece of data, creating channels to "share your memory" is overkill.

# Race Detection

What's a data race?  When concurrent goroutines try to access the same memory, with at least one of them being a mutate/write.

By far, data races are the hardest to debug problems you'll come across. Why?

- Manifestation of bugs aren't deterministic. They depend on the state of multiple goroutines.
- Attempting to diagnose the bug only furthers the race - fmt.Println(someValue) - it's another access!

Fortunately, Go has built-in race detection tools.  The `go` command has a `-race` flag that can be applied to:

- go test
- go build
- go install


## Race Example

	cd $GOPATH/src/github.com/gophertrain/material/wranglingconcurrency/demos/race/

## Race Exercise

Run the race example in `$GOPATH/src/github.com/gophertrain/material/wranglingconcurrency/demos/race/`

	go run main.go

What happens?

Now build with the `-race` flag.

	go build  -race .
	./race

What results do you get?

## Race Exercise - Individual Exercise

Fix the race example in `$GOPATH/src/github.com/gophertrain/material/wranglingconcurrency/exercises/race/` using tools you have learned in this module. Test it by running with the `-race` flag.

	Never ship a binary built with `-race`.
	It will use many times more memory and be significantly slower.

# Conclusion

Go's concurrency features are among the easiest to use in any programming language.  With great power comes great responsibility.  

- Learn Channel behavior. All of it.
- Use sync.WaitGroup and sync.ErrGroup to make your life easier
- Use mutexes for simple protection
- Use the race detector locally when you build and test your code. (But not in the final binary) 
- Add -race to your `go test` scripts, especially in your CI pipeline.
