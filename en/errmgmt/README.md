# Effective Error Management

In this section you’ll learn how to create and handle errors effectively.  

Topics:

- Creating Errors
- Handling Errors
- Differences between libraries and applications
- Guidelines

You’ll end up with a clear guideline for error handling, in both your library code and your applications.

# Philosophy: Go and Errors

Proverb:

	Errors are just values.
_Rob_Pike_ at [Gopherfest SV, 2015](https://www.youtube.com/watch?v=PAAkCSZUG1c)

[Other Go Proverbs](https://go-proverbs.github.io/)

### What does that mean?

Functionally, it means that *error* is an interface in Go.  Any type that implements the function 

	Error() string

can be used as an *error* type.

Conceptually, it means that you have a lot of freedom and flexibility to do error handling.

# Creating Errors

## Creating Errors - errors.New()

You can use the `errors` package to create a descriptive error on the fly

	return errors.New("Login Failed")

- Simple
- Least amount flexibility

## Creating Errors - Wrapping an Error

You can use fmt.Errorf() to wrap an error with more context

	return fmt.Errorf("Login failed:", err)

- Adds some context
- Harder to determine original error type


## Creating Errors - Exporting Named Errors

You can create known error types and return them 

	var (
		// ErrInvalidLogin is returned when a bad username or password
		// is used to attempt a login to the application
		ErrInvalidLogin = errors.New("Invalid Login")
		// ErrDBUnavailable is returned with the database can't be contacted
		ErrDBUnavailable = errors.New("Database Unavailable")
	)

	// func DoSomething() {
		// some code

		err := user.Login(u, p)
		if err != nil {
			return ErrInvalidLogin
		}

		// normal execution
	}

- Easy to compare and test
- Requires discipline to remember to define these when you find new error conditions

## Creating Errors - Exporting Custom Types 

You can create a custom error type	

	type ErrCode int

	var (
		LoginFailed ErrCode = 1000
	)

	var appErrors = map[ErrCode]string{
		LoginFailed: "Login Failed",
	}

	type ApplicationError struct {
		Time time.Time
		Code ErrCode
	}

	func (e *ApplicationError) Error() string {
		return fmt.Sprintf("%d: %s", e.Code, appErrors[e.Code])
	}
	func NewError(code ErrCode) *ApplicationError {
		return &ApplicationError{
			Time: time.Now(),
			Code: code,
		}
	} 

- Great flexibility
- Good when you might want additional error state stored in the error (time)
- I've used these in web applications/API's where I want to send a very consistent message back to the user, serialized as JSON.


## Creating Errors - Your Own Types

You can implement the `error` interface on your domain types

	type Customer struct {
		ID     int
		Name   string
		ErrMsg string
	}

	func (c *Customer) Error() string {
		if c.Error != nil {
			return fmt.Sprintf("Error: %s", c.ErrMsg)
		}
		return nil
	} 

- Creative
- Requires setting your type's ErrMsg equivalent when an error occurs
- Great for a state machine

# Handling Errors

There are really three ways to test error conditions

- Compare with known values
- Assert error type
- Ignore and pass up the call stack

## Handling Errors - Sentinel Errors 

If an error is your own package, or exported from an external package, you can test the error in your code.

Dave Cheney calls these "Sentinel errors" *[1]

	if err == io.EOF


[1] http://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully


	func Login(u, p string) error {
		available := db.Alive() //bool
		if !available {
			return ErrDatabaseUnavailable
		}
		err := user.Login(u, p)
		if err != nil {
			return ErrInvalidLogin
		}
		return nil
	}

	func main() {
		err := Login(os.Args[1], os.Args[2])
		if err != nil {
			if err == ErrDatabaseUnavailable {
				fmt.Println("Database Not Available, try later")
			}
			if err == ErrInvalidLogin {
				fmt.Println("Bad Login")
			}
			os.Exit(-1)
		}
	} 

- No context available to the caller, just a known error condition
- High coupling between caller and called packages
- Defeated if you add extra context by wrapping the error:
	return fmt.Errorf("DB Error:", ErrDBUnavailable)
	// NOT == ErrDBUnavailable now

- Many packages export error types, especially in network, OS, and database operations

## Handling Errors - Assert with Custom Type

Using custom types provides you with more flexibility in your error handling logic.

	func Login(u, p string) (*Customer, error) {
		cust, err := user.Login(u, p)
		if err != nil {
			return nil, NewError(LoginFailed)
		}
		return cust, nil
	}

	func main() {
		cust, err := Login(os.Args[1], os.Args[2])
		switch err := err.(type) {
		case nil:
			// OK to proceed
			ProcessTasks()
		case *ApplicationError:
			log.Println("%v:", err.Time, err.Error())
		default:
			// stop execution - unknown error
			log.Fatal(err)
		}
	}
	
- Testing for error type lets you embed useful context in your error
- Error handling can be very specific
- Still highly coupled, sometimes this is OK and preferred

## Handling Errors - Ignore and Pass Up The Stack

	return db.Login(u,p)

- Not My Problem(tm)
- Do this deep internally, not in exported Functions/Methods

## Libraries are different

Internal error handling vs external library API

- Libraries should export known error types for conditions that are static

	sql.ErrNoRows
	io.EOF
	bytes.ErrTooLarge 

- Return underlying error if there's nothing the library can do to fix things

	Network calls fail - you can't fix this, return the error
	ex: log/syslog#Dial()

- Wrap error 
	
	NO.
	return fmt.Errorf("reading from database:", err)
	// can't test for known error without strings.Contains() 
	// Destroyed original context
	


## Handling Errors - General Advice

### Only handle errors once - NO

	func GetConfig() error {
		bb, err := ioutil.ReadFile("/etc/app/config")
		if err != nil {
			log.Println("Couldn't read config:", err)
			return nil, err
		}
		return bb, nil
	}	

If the caller logs the error too, you've got possibly two log lines with confusingly similar messages.

### Only handle errors once - BETTER

	func GetConfig() ([]byte, error) {
		return ioutil.ReadFile("/etc/app/config")
	}

Since you're not doing anything to resolve the error in your function, return it unchanged.


### Wrapping Errors For Clarity

In early 2016, Dave Cheney released *github.com/pkg/errors*

This package allows you to have the best of all worlds, by letting you wrap errors in contextual information while preserving the original error type.

	// open config file
	_, err := ioutil.ReadAll(r)
		if err != nil {
		return errors.Wrap(err, "read failed")
	}

Callers that print the error will get the wrapped error with your prefix:

	read failed: file not found


Every function in the stack can wrap the error with their own context:

	return errors.Wrap(err, "opening configuration file")

	// opening configuration file: read failed: file not found
	
You can retrieve the original cause of the error by using the errors.Cause() function:

	originalError := errors.Cause(err)


### Wrapping Errors For Clarity - example

	func getConfig(path string) ([]byte, error) {
		bb, err := ioutil.ReadFile(path)
		return bb, errors.Wrap(err, "get config")
	}

	func initialize(path string) error {
		config, err := getConfig(path)
		if err != nil {
			return errors.Wrap(err, "initialize")
		} // process the config file, open database, etc
	}

	func main() {
		if err := initialize(os.Args[1]); err != nil {
			errors.Print(err)
			switch err := errors.Cause(err).(type) {
			case os.PathError:
				log.Error("Unable to open log file, check location:", err)
			default:
				log.Error("Unknown error:", err)
			}
			os.Exit(-1)
		}
	}

github.com/pkg/errors is under consideration for inclusion in the standard library.  (which is rare!)

Use it.


# Exercise

Discuss pros and cons of error handling.

- Do you have examples of projects that do it well?
- Do you have examples of confusing projects?


