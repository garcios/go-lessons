## About this module
The common and considered best practice for separating interfaces is to align them with the consuming package or type 
that uses the interface, rather than the package or type that implements the interface.

Here are some details on this practice:
1. Interfaces in the consumer package/type: The main idea here is to define an interface right where you need it 
(where you're going to use it). This way, your interface is only as broad as required by the consumer. 
You are not required to create large interfaces with many methods that may not be used in some implementations. 
This practice also helps keep your code organized and easy to understand. An interface grouped with the consumer type 
can look something like this:

```go
package consumer
    
type MyInterface interface{
    Function1()
    // Function2()
    // Add more functions as needed
}
```
2. Minimal interface: Instead of creating large monolithic interfaces, Go encourages the creation of minimal interfaces 
(i.e., interfaces with the smallest possible number of methods). This again aligns with the point above and follows the 
Go proverb of "The bigger the interface, the weaker the abstraction."

3. Multiple small files vs. One large file: In Go, it is common to separate different things into different files. 
This can be different types, different interfaces, different functions, etc. It can be useful to separate interfaces 
into their own files named for what the interface represents. The package tree can give you a good idea of what 
interfaces are available without having to open a single huge file.

## Checking interface compliance

In Go, unlike some other languages such as Java or C#, there's no explicit declaration of implements keyword that 
connects an interface with its implementing struct. Instead, you can just create a struct with methods that match the 
interface signature - and the struct will implicitly satisfy the interface.

The given line of code.
```go
var _ Fetcher = (*Gryffindor)(nil)
```
is checking at compile-time that the type *Gryffindor satisfies the Fetcher interface. If *Gryffindor does not satisfy 
the Fetcher interface, the code will not compile.

This declaration does not create a usable variable and underscore _ is a write-only variable in Go that's used when 
syntax requires a variable name but program logic does not.

So, the above code doesn't do anything functional at runtime, but it helps you to ensure correct code at compile time.

As for the provided code snippets:

1. Fetcher interface is defined with two methods, GetAll() and GetSourceName().
2. Gryffindor struct is defined and it implements GetSourceName and GetAll methods, plus two additional methods 
   GetByID and GetHouseHead.

Noting, that these additional functions do not disrupt the fact that Gryffindor is a Fetcher.

Remember that in Go, a type does not have to implement all (and only) the methods of an interface to satisfy it, 
it needs to implement at least those, hence the Gryffindor structure still satisfies the Fetcher interface.

## Instructions
```shell
go mod init github.com/oskiegarcia/fetcher
```
### 2. Creat main.go

### 3. Run the program
```shell
go run .
```

