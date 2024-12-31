## About this module
This is a banking application that handles the following functionalities:
- Check account balance
- Deposit amount
- Withdraw amount

## Error handling
In general, if a function returns multiple values, the last value will be an error object.
Error is handled by checking whether the returned err value is not nil.

e.g.
```go

func main() 
{
  account, err := getAccount()
  if err != nil {
    fmt.Printf("Error: %v", err)
  }
	
}

func getAccount() (*Account, error)
{
	return nil, errors.New("An error occurred!")
}

```
### fmt.Errorf vs errors.New
Both fmt.Errorf and errors.New are used to create a new error, but they are used in slightly different scenarios.

1. errors.New: This function is used when you want to create a simple error message without any formatting or additional arguments. It just takes a string message and returns a new error with that message.
Here is an example of using errors.New:

```go
err := errors.New("This is a simple error")
```

2. fmt.Errorf: This function is used when you want to create a new error with formatted message. It works like 
fmt.Printf but returns an error instead of printing to standard output. You can use this function when you want to 
include values in your error message.

Here is an example of using fmt.Errorf:
```go
err := fmt.Errorf("Error happened at %v", time.Now())
```
In conclusion:
- Use errors.New when you just need a simple, static error message.
- Use fmt.Errorf when you need to include extra information in your error message (like variables or any other dynamic
parts).

### Other error types
1. Custom Error Types

You can define your own error type by creating a type that implements the error interface. The error interface 
requires an Error() string method.

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("myerror: code %v - %s", e.Code, e.Message)
}
```

This allows you to add more context to the error beyond a simple string message.

2. The pkg/errors Package

The package github.com/pkg/errors provides additional functions to annotate and wrap errors. This provides a stack 
trace and the line number where the error occurred, which can be invaluable when debugging.

```go
import "github.com/pkg/errors"

...

err := errors.Wrap(err, "read failed")
```

3. Sentinel Errors

Sentinel errors are predefined errors that you can compare against, usually defined as constants. For example, the 
io package defines a sentinel error io.EOF to indicate that the end of a file or stream has been reached.

```go
if err == io.EOF {
    fmt.Println("Reached the end of the file")
}
```

4. Error Wrapping (From Go 1.13)
Starting from Go 1.13, the standard fmt package supports a method for wrapping errors using the %w verb. 
This can be used in conjunction with errors.Is and errors.As methods to check and extract wrapped errors.

```go
originalErr := someFunction()
return fmt.Errorf("operation failed: %w", originalErr)

// Later or elsewhere...
if errors.Is(err, originalErr) {
    // handle specific error case
}
```

Wrapping errors allows you to add context information to the error while still being able to check if the error is a 
specific error. In the above example, originalErr is wrapped with additional context information "operation failed: ",
but you can still use errors.Is to check if an error is originalErr.

### Panic function

The panic function in Go language is a built-in function that is used to abort the flow of control and start panicking.
When the panic function is called, the normal execution of the program is stopped and the function starts to collapse 
or panic which defers the functions that are in the stack. Basically, it's Go's mechanism for handling unexpected or 
unrecoverable situations, somewhat analogous to throwing exceptions in other programming languages.

Here is the function signature:
```go
func panic(v interface{})
```

The panic function accepts any value -- we can provide an object of any type.
Here is an example of how to use panic:

```go
func main() {
    fmt.Println("Start")
    panic("Something bad happened")
    fmt.Println("End")  // This will not be reached
}
```

In this example, the program will print "Start", then it will panic with the message "Something bad happened". The rest
of the code following the panic invocation will not be executed.
One of the important things to know about using panic is that any deferred functions will be executed before the 
program crashes. This is very useful for cleanup activities. However, in general, it's considered good practice to use 
panic judiciously and resort to error handling mechanisms for most error situations.

### Recover function
The recover function is used to regain control of a panicking goroutine.

recover is a built-in function that regains control of a panicking goroutine. When called within a deferred function, 
recover stops the panicking sequence and restores normal execution.

Here is the function signature:
```go
func recover() interface{}
```

If recover is called directly as part of the normal execution (not in deferred function or after panic sequence has 
started), it doesn't do anything and returns nil.

he standard way to use recover is in a deferred function within the same goroutine that's executing the panic sequence. 
Here's an example:
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from", r)
        }
    }()
    
    fmt.Println("Start")
    panic("Something bad happened")
    fmt.Println("End")  // This will not still be reached because of the panic
}
```

In this example, the recover function is being called within a deferred function. If a panic happened somewhere in the 
code (in this case, just one line after the defer), the deferred function would be executed during the panic sequence, 
recover would stop the panic, executing immediately after the call to panic. After recovery, the program terminates 
normally and the "End" line is still not reached since it's unreachable after a panic.

It's important to note that recovering from a panic is not the same as handling an error. A panic generally indicates a
major fault in the program that cannot be easily recovered without risking further errors, and is usually used for 
preventing the program from running into a possibly invalid state. Errors, on the other hand, are a more graceful way of
handling and notifying about abnormal situations in the program flow.
