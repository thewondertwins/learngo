#  Writing Scalable Services

In this section we’ll demonstrate advanced patterns and techniques for building applications that scale beyond your staging environment and into real-world usage.  You’ll see how to prevent resource exhaustion by using timeouts and cancellations.  You’ll learn how to measure the potential performance of your application and determine the theoretical maximums that your application can handle with a single instance.  You’ll have labs and exercises that show you how to measure and monitor your application while it’s running.

- Timeouts and Cancellations using Contexts
- Measuring Throughput
- Determining Your Limits
- Measuring, Monitoring, Logging


# Timeouts and Cancellations using Contexts

Internal policy at Google is that all functions accept a context as the first parameter.

Why?

- Cancellation
- Deadlines
- Timeouts
- Context scoped k/v retrieval & storage

## Cancellation

[GoDoc](https://golang.org/pkg/context/#WithCancel)

	cd $GOPATH/src/github.com/gophertrain/material/scalablesvcs/demos/contexts/cancellation/
	go run main.go

## Cancellation Walkthrough

	type Result struct {
		Value int
		Err   error
	}

	func main() {
		// launch 10k goroutines, trying to get random
		// to return "150"
		result := Random(context.Background(), 150, 10000)
		if result.Err != nil {
			fmt.Println("error:", result.Err)
		}
		fmt.Println(result.Value)
	}

	func Random(ctx context.Context, target, count int) Result {
		rand.Seed(82)
		c := make(chan Result, count)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		search := func(target int) {
			sleep := time.Duration(rand.Intn(10))
			select {
			case <-ctx.Done():
				return
			case <-time.After(sleep):
				r := rand.Intn(1000)
				if r == target {
					c <- Result{Value: r}
					return
				}
				fmt.Printf(".")
			}
		}
		for i := 0; i < count; i++ {
			go search(target)
		}
		select {
		case <-ctx.Done():
			return Result{Err: ctx.Err()}
		case r := <-c:
			return r
		}

	} 

## Cancellation BUG

What happens if 10000 goroutines don't return the target number?
Odds are pretty good in our favor that they will, but it could happen.

What will be the result?

Deadlock! (forced by changing 10k to just 10)
	
	..........fatal error: all goroutines are asleep - deadlock!
	goroutine 1 [select]:
	main.Random(0x111480, 0xc420084000, 0x96, 0xa, 0x0, 0x0, 0x0)
	/Users/bketelsen/src/github.com/gophertrain/modules/src/scalablesvcs/demos/contexts/cancellation/main.go:45 +0x308
	main.main()
	/Users/bketelsen/src/github.com/gophertrain/modules/src/scalablesvcs/demos/contexts/cancellation/main.go:18 +0x4f

## Deadline

You can fix this by providing a deadline.  Deadlines are propagated through the call chain, and are absolute.  A deadline is a fixed point in time. 

```
When time(X) arrives, call the cancelFunc
```

	type Result struct {
		Value int
		Err   error
	}

	func main() {
		// launch 10k goroutines, trying to get random
		// to return "150"
		// derive context with dedaline set to 5 seconds
		deadline, cf := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
		defer cf()

		// passing context with deadline method will be terminated if not completed within 5s
		result := Random(deadline, 150, 10000)
		if result.Err != nil {
			fmt.Println("error:", result.Err)
		}
		fmt.Println(result.Value)

	}

	func Random(ctx context.Context, target, count int) Result {
		rand.Seed(82)
		c := make(chan Result, count)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		search := func(target int) {
			sleep := time.Duration(rand.Intn(10))
			select {
			case <-ctx.Done():
				return
			case <-time.After(sleep):
				r := rand.Intn(1000)
				if r == target {
					c <- Result{Value: r}
					return
				}
				fmt.Printf(".")
			}
		}
		for i := 0; i < count; i++ {
			go search(target)
		}
		select {
		case <-ctx.Done():
			return Result{Err: ctx.Err()}
		case r := <-c:
			return r
		}

	} 


## Timeout

	type Result struct {
		Value int
		Err   error
	}

	func main() {
		// launch 10k goroutines, trying to get random
		// to return "150"
		// Timeout is relative
		timeout, cf := context.WithTimeout(context.Background(), 5*time.Second)
		defer cf()
		result := Random(timeout, 150, 10)
		if result.Err != nil {
			fmt.Println("error:", result.Err)
		}
		fmt.Println(result.Value)

	}

	func Random(ctx context.Context, target, count int) Result {
		rand.Seed(82)
		c := make(chan Result, count)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		search := func(target int) {
			sleep := time.Duration(rand.Intn(10))
			select {
			case <-ctx.Done():
				return
			case <-time.After(sleep):
				r := rand.Intn(1000)
				if r == target {
					c <- Result{Value: r}
					return
				}
				fmt.Printf(".")
			}
		}
		for i := 0; i < count; i++ {
			go search(target)
		}
		select {
		case <-ctx.Done():
			return Result{Err: ctx.Err()}
		case r := <-c:
			return r
		}

	} 


## Timeout vs. Deadline

There is no practical difference between timeout and deadline.  One is a relative time, and one is an absolute time.

In fact, from godoc:

	WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

Use the one that feels comfortable to you.

## Exercise

	src/scalablesvcs/demos/contexts

Exercise:
- Open each of the demos in the contexts directory.
- Change the values of the loops and timeouts and experiment with the results.

## Context Values

You can use a context.Context as a key/value store to store and retrieve *request*specific* information.

- Don't store dependencies
- Don't store your logger, your database handle, etc

Store request/instance specific information
- RequestID
- Server processing initial request
- Originating IP
- Trace ID
- Authentication Token

Never depend on a value being present

Always use a default if the value isn't present, or skip that particular logic

Use typed keys to avoid collision:

	type key int

	var RequestID key = 0 
	ctx = context.WithValue(ctx, RequestID, "12345")


	reqid, ok := ctx.Value(RequestID) // will work

	var intID int = 0
	reqid, ok := ctx.Value(RequestID) // will not work

Now keys are scoped to the package.  Even though the "value" of the key is an integer with the value "0", it's a strongly typed value.  In this case `values.Main#RequestID` which will not collide with any other key even if the underlying integer value is also "0".

# Measuring Throughput

Measuring your throughput is as subjective a topic as we can cover.  Instead of concrete examples, we'll cover concepts and plans.


## Measuring Throughput

Start with micro benchmarks.  Worry about things with complex CALCULATIONS first.  

- Benchmark functions with mocks, determine how long a particular function takes with fake data source
- Benchmark with real dependencies, determine how long it takes with staging/production data source

Determine overhead without extra latency first, then with the latency brought by your data source.

- Guard these with a build flag:
	// +build integration
	// This way these benchmarks can/will only run in your integration environment.
- Output the benchmarks using your CI environment to a timestamped file.
- Compare benchmarks against previous results to prevent performance regressions

This will give you a baseline of what the function costs and what the function costs with the external dependency.

### Throughput Demo

	cd $GOPATH/src/github.com/gophertrain/material/scalablesvcs/demos/throughput

### Throughput in CI


	cd $GOPATH/src/github.com/gophertrain/material/scalablesvcs/demos/throughput/Makefile

	$ make compare	
	 
Will fail if the delta is positive.  (Hacky implementation)

Example implementation - won't work as-is.  You need to check in the previous results into source control in the CI process to do this.


# Determining Your Limits

After measuring throughput of individual components, you need to know how to determine your application's limits.

Again, there is no hard and fast rule to this, only general steps.


### Step 1:

Stand up your application in an environment as similar to production as possible.

Anything that's different from production is a variance that decreases your trust in your measurements.

### Step 2:

Generate load to test the application.  Measure performance.

- Generate realistic load.  Don't fetch the same data from sources that may cache it, for example.
- Use more than one client to generate load. This reduces network limitations, and increases concurrency contention.

### Step 3:

- Use the profiling tools from module 6 to investigate performance bottlenecks
- /debug/pprof *every time*  - use the builtin performance capturing tools!


# Measuring, Monitoring, Logging

##  Measuring

You should capture relevant metrics for every [significant] function.

You determine what "significant" means.  Probably not string manipulation, definitely anything that touches disk, network, etc.

There are dozens of ways to capture metrics. `Prometheus` is the best way to capture, collect and process them.

My *strong* opinion:  

- if you're unsure how you'll be monitoring them, use github.com/armon/go-metrics
- if you know you can use Prometheus, use the Go prometheus library directly

armon/go-metrics will export to statsd, prometheus and others.

## Monitoring

Monitor externally from your application.  

- prometheus alertmanager
- datadog
- grafana alerts
- etc

## Logging

- Always use structured logging

Recommendation:

	github.com/uber-go/zap

- FAST
- Very low memory overhead / low allocations
- HTTP Handler to allow you to change log level at RUNTIME
- JSON or text output.  Use a flag to choose which level based on whether you're running locally or in production

## Measuring & Logging

Recommendation:

Do these together. Wrap one of these:
- https://github.com/golang/net/tree/master/trace 
- https://github.com/sourcegraph/appdash
- https://github.com/opentracing/basictracer-go

and include logging.

https://github.com/bketelsen/trace
	
	see @rakyll's new package - getting better

# Wrapup

Your services won't scale without paying attention to all the details.

- Measure
- Monitor



