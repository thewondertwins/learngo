# Writing Code To Be Fast The First Time

In this section we’ll build on our learnings from profiling and benchmarking and show advanced techniques to create code that is fast the first time. You’ll learn how to minimize the garbage created in your application, how to build reusable pools of expensive resources, and other advanced tricks to make your applications faster and avoid problems before they happen.

- Arrays and Slices
- Make GC Faster by Creating Less Garbage
- Pools
- Avoiding Unnecessary Conversions
- Use the Stack

# Memory management and GC tuning

Go is a garbage collected language. This is a design principle, it will not change.

As a garbage collected language, the performance of Go programs is often determined by their interaction with the garbage collector.

The Go GC favors lower latency over maximum throughput; it moves some of the allocation cost to the mutator to reduce the cost of cleanup later.

Next to your choice of algorithms, memory consumption is the most important factor that determines the performance and scalability of your application.

This section discusses the operation of the garbage collector, how to measure the memory usage of your program and strategies for lowering memory usage if garbage collector performance is a bottleneck.

## Garbage collector design

The design of the Go GC has changed over the years

- Go 1.0: stop the world mark sweep collector based heavily on tcmalloc.
- Go 1.3: fully precise collector, wouldn't mistake big numbers on the heap for pointers.
- Go 1.5: new GC design, focusing on _latency_ over _throughput_.
- Go 1.6: GC improvements, handling larger heaps with lower latency.
- Go 1.7: small GC improvements, mainly refactoring.
- Go 1.8: ROC collector 

## Garbage collector world view

	The purpose of a garbage collector is to present the illusion that there is an infinite amount of memory.

You may disagree with this statement, but this is the base assumption of how garbage collector designers think.

A *stop the world, mark sweep* GC is the most efficient in terms of total run time; good for batch processing, simulation, etc.

The Go GC favors *lower_latency* over *maximum_throughput*; it moves some of the allocation cost to the mutator to reduce the cost of cleanup later.

The Go GC is designed for low latency servers and interactive applications.

## Garbage collector monitoring

A simple way to obtain a general idea of how hard the garbage collector is working is to enable the output of GC logging.

These stats are always collected, but normally suppressed, you can enable their display by setting the `GODEBUG` environment variable.

	% env GODEBUG=gctrace=1 godoc -http=:8080
	gc 1 @0.017s 8%: 0.021+3.2+0.10+0.15+0.86 ms clock, 0.043+3.2+0+2.2/0.002/0.009+1.7 ms cpu, 5->6->1 MB, 4 MB goal, 4 P
	gc 2 @0.026s 12%: 0.11+4.9+0.12+1.6+0.54 ms clock, 0.23+4.9+0+3.0/0.50/0+1.0 ms cpu, 4->6->3 MB, 6 MB goal, 4 P
	gc 3 @0.035s 14%: 0.031+3.3+0.76+0.17+0.28 ms clock, 0.093+3.3+0+2.7/0.012/0+0.84 ms cpu, 4->5->3 MB, 3 MB goal, 4 P
	gc 4 @0.042s 17%: 0.067+5.1+0.15+0.29+0.95 ms clock, 0.20+5.1+0+3.0/0/0.070+2.8 ms cpu, 4->5->4 MB, 4 MB goal, 4 P
	gc 5 @0.051s 21%: 0.029+5.6+0.33+0.62+1.5 ms clock, 0.11+5.6+0+3.3/0.006/0.002+6.0 ms cpu, 5->6->4 MB, 5 MB goal, 4 P
	gc 6 @0.061s 23%: 0.080+7.6+0.17+0.22+0.45 ms clock, 0.32+7.6+0+5.4/0.001/0.11+1.8 ms cpu, 6->6->5 MB, 7 MB goal, 4 P
	gc 7 @0.071s 25%: 0.59+5.9+0.017+0.15+0.96 ms clock, 2.3+5.9+0+3.8/0.004/0.042+3.8 ms cpu, 6->8->6 MB, 8 MB goal, 4 P

The trace output gives a general measure of GC activity.

Using `GODEBUG=gctrace=1` is good when you *know* there is a problem, but for general telemetry on your Go application I recommend the `net/http/pprof` interface.

    import _ "net/http/pprof"

Importing the `net/http/pprof` package will register a handler at `/debug/pprof` with various runtime metrics, including:

- A list of all the running goroutines, `/debug/pprof/heap?debug=1`. 
- A report on the memory allocation statistics, `/debug/pprof/heap?debug=1`.

*Warning*: `net/http/pprof` will register itself with your default `http.ServeMux`.

Be careful as this will be visible if you use `http.ListenAndServe(address, nil)`.

## Garbage collector tuning

The Go runtime provides one environment variable to tune the GC, `GOGC`.

The formula for GOGC is as follows.

    goal = reachable * (1 + GOGC/100)

For example, if we currently have a 256mb heap, and `GOGC=100` (the default), when the heap fills up it will grow to

    512mb = 256mb * (1 + 100/100)

- Values of `GOGC` greater than 100 causes the heap to grow faster, reducing the pressure on the GC.
- Values of `GOGC` less than 100 cause the heap to grow slowly, increasing the pressure on the GC.

The default value of 100 is *just a guide*. you should choose your own value *after profiling your application with production loads*.

# Reduce allocations

Make sure your APIs allow the caller to reduce the amount of garbage generated.

Consider these two Read methods

    func (r *Reader) Read() ([]byte, error)
    func (r *Reader) Read(buf []byte) (int, error)

The first Read method takes no arguments and returns some data as a `[]byte`. The second takes a `[]byte` buffer and returns the amount of bytes read.

The first Read method will **always** allocate a buffer, putting pressure on the GC. The second fills the buffer it was given.

[ReadAtLeast](https://golang.org/pkg/io/#ReadAtLeast)

[ReadFull](https://golang.org/pkg/io/#ReadFull)

In this example, `Conn.Loop` can run forever without generating garbage.

	type Conn struct {
		r  io.ReadCloser
		ch chan uint32
	}

	func (c *Conn) Loop() {
		defer r.Close()
		var buf [512]byte
		for {
			b := buf[:] // create slice of buf
			n, err := c.r.Read(b)

			for b = b[:n]; len(b) != 0; b = b[4:] {
				ch <- binary.BigEndian.Uint32(b)
			}

			if err != nil {
				return
			}
		}
	}

## strings and []bytes

In Go `string` values are immutable, `[]byte` are mutable.

Most programs prefer to work `string`, but most IO is done with `[]byte`.

Avoid `[]byte` to string conversions wherever possible, this normally means picking one representation, either a `string` or a `[]byte` for a value. Often this will be `[]byte` if you read the data from the network or disk.

The [bytes](https://golang.org/pkg/bytes/) package contains many of the same operations—`Split`, `Compare`, `HasPrefix`, `Trim`, etc—as the [strings](https://golang.org/pkg/strings/) package.

Under the hood `strings` uses same assembly primitives as the `bytes` package.

## Using []byte as a map key

It is very common to use a `string` as a map key, but often you have a `[]byte`.

The compiler implements a specific optimisation for this case

     var m map[string]string
     v, ok := m[string(bytes)]

This will avoid the conversion of the byte slice to a string for the map lookup. This optimisation is very specific, it won't work if you do something like

     key := string(bytes)
     val, ok := m[key] 

## Avoid string concatenation

Go strings are immutable. Concatenating two strings generates a third.

Avoid string concatenation by appending into a `[]byte` buffer.

_Before:_

    s := request.ID
    s += " " + client.Address().String()
    s += " " + time.Now().String()
    return s

_After:_

    b := make([]byte, 0, 40) // guess
    b = append(b, request.ID...)
    b = append(b, ' ')
    b = append(b, addr.String()...)
    b = append(b, ' ')
    b = time.Now().AppendFormat(b, "2006-01-02 15:04:05.999999999 -0700 MST")
    return string(b)

## Preallocate slices if the length is known

Append is convenient, but wasteful.

Slices grow by doubling up to 1024 elements, then by approximately 25% after that. What is the capacity of `b` after we append one more item to it?

	func main() {
		b := make([]int, 1024)
		b = append(b, 99)
		fmt.Println("len:", len(b), "cap:", cap(b))
	} 
	
	Output: len: 1025 cap: 1344

If you use the append pattern you could be copying a lot of data and creating a lot of garbage.

If know know the length of the slice beforehand, then pre-allocate the target to avoid copying and to make sure the target is exactly the right size. 

_Before:_

     var s []string
     for _, v := range fn() {
            s = append(s, v)
     }
     return s

_After:_

     vals := fn()
     s := make([]string, len(vals))
     for i, v := range vals {
            s[i] = v           
     }
     return s

## Slice Tidbit

Slice Optimization:

	//s := []string{} // not this
	var s []string // this!
	for t := range things {
		s = append(s, t)
	}

Creating `s` without initializing it will never allocate memory if there is nothing in the `things` collection.

Creating `s` with an initialization will *always* allocate memory.  Optimize for the best case and use the `var s []string` form.

## Using sync.Pool

The `sync` package comes with a `sync.Pool` type which is used to reuse common objects.

`sync.Pool` has no fixed size or maximum capacity. You add to it and take from it until a GC happens, then it is emptied unconditionally. 

	var pool = sync.Pool{
		New: func() interface{} { return make([]byte, 4096) },
	}

	func fn() {
		buf := pool.Get().([]byte) // takes from pool or calls New
		// do work
		pool.Put(buf) // returns buf to the pool
	} 

*Warning*: `sync.Pool` is not a cache. It can and will be emptied *at any time*.

Do not place important items in a `sync.Pool`, they will be discarded.

## Rolling Your Own Pool

It's easier to make a safe pool using buffered channels.

Start with a type that represents your pool:

	type Pool struct {
		pool chan *Thing
	}

Then create a New() function that creates a new pool

	// NewPool creates a new pool of Things
	func NewPool(max int) *Pool {
		return &Pool{
			pool: make(chan *Thing, max),
		}
	} 

The NewPool() function takes an integer parameter that represents the maximum size of the pool.  That way the caller can control the size of the pool.


To take an object from the pool, use Borrow():

	// Borrow a Thing from the pool.
	func (p *Pool) Borrow() *Thing {
		var t *Thing
		select {
		case t = <-p.pool:
		default:
			t = newThing()
		}
		return t
	} 

Borrow uses all the properties of channels that we love.  It tries to read a from the channel, but creates a new object if there's nothing to read.

To return an object to the pool, use Return(), preferably in a `defer`:

	func (p *Pool) Return(t *Thing) {
		select {
		case p.pool <- t:
		default:
		}
	}

Return tries to send the object back on the channel, and discards it if it can't, because the pool is at max capacity.

Where does this make sense?  

- Things that have costly setup, or large memory footprint
- Encoders/Decoders with their own internal buffer
- Rate-limiting the number of outbound connections to an external service

## Use the Stack, Luke

Memory that is allocated on the heap is more costly to collect.  Memory that is allocated on the stack is less costly and easier to collect.

	In the current compilers, if a variable has its address taken, that variable is 
	a candidate for allocation on the heap. However, a basic escape analysis 
	recognizes some cases when such variables will not live past the return from 
	the function and can reside on the stack.

[Golang FAQ](https://golang.org/doc/faq)

You can help keep things on the stack by avoiding making pointers when they aren't needed. It's less pressure on the garbage collector to pass a value to a function *if the fuction won't be mutating the value*.

You can't control this 100%. But you can sway things in your favor by avoiding pointers unless they're needed.

You can see what the compiler will do in a specific instance by calling the `go` compiler with `-gcflags '-m'` which will print verbose output of escape analysis information.

Example of output:

	./gogrep.go:57: inlining call to os.FileMode.IsRegular
	./gogrep.go:60: inlining call to strings.HasSuffix
	./gogrep.go:28: inlining call to flag.NArg
	./gogrep.go:32: inlining call to flag.Arg
	./gogrep.go:32: inlining call to Arg
	./gogrep.go:33: inlining call to flag.Arg
	./gogrep.go:33: inlining call to Arg
	./gogrep.go:34: inlining call to "golang.org/x/net/context".Background
	./gogrep.go:45: leaking param: ctx
	./gogrep.go:50: func literal escapes to heap
	./gogrep.go:50: func literal escapes to heap
	./gogrep.go:47: make(chan string, 100) escapes to heap
	./gogrep.go:45: leaking param: root
	./gogrep.go:77: func literal escapes to heap
	./gogrep.go:77: func literal escapes to heap
	./gogrep.go:45: leaking param: pattern
	./gogrep.go:74: make(chan string, 100) escapes to heap
	./gogrep.go:93: func literal escapes to heap
	./gogrep.go:93: func literal escapes to heap
	./gogrep.go:53: leaking closure reference root
	./gogrep.go:53: func literal escapes to heap
	./gogrep.go:53: leaking param: info
	./gogrep.go:53: leaking param: path
	./gogrep.go:66: leaking closure reference ctx
	./gogrep.go:66: leaking closure reference ctx
	./gogrep.go:78: leaking closure reference p
	./gogrep.go:78: leaking closure reference p
	./gogrep.go:87: leaking closure reference ctx
	./gogrep.go:87: leaking closure reference ctx
	./gogrep.go:94: leaking closure reference g
	./gogrep.go:53: leaking param: err to result ~r3 level=0
	./gogrep.go:82: search.func2 ([]byte)(pattern) does not escape
	./gogrep.go:20: func literal escapes to heap
	./gogrep.go:20: func literal escapes to heap
	./gogrep.go:34: "golang.org/x/net/context".Context("golang.org/x/net/context".background) escapes to heap
	./gogrep.go:37: err escapes to heap
	./gogrep.go:40: name escapes to heap
	./gogrep.go:42: len(m) escapes to heap
	./gogrep.go:42: "hits" escapes to heap
	./gogrep.go:21: os.Args[0] escapes to heap
	./gogrep.go:22: "Usage:" escapes to heap
	./gogrep.go:24: "Flags:" escapes to heap
	./gogrep.go:37: main ... argument does not escape
	./gogrep.go:40: main ... argument does not escape
	./gogrep.go:42: main ... argument does not escape
	./gogrep.go:21: main.func1 ... argument does not escape
	./gogrep.go:22: main.func1 ... argument does not escape
	./gogrep.go:24: main.func1 ... argument does not escape 

# Exercises

Run go test -bench=. and benchmark the various concatenation strategies to see the difference between string manipulation and byte buffers.

