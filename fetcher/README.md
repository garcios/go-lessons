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
(i.e., interfaces with the smallest possible number of methods). This again aligns with the point above and follows the Go proverb of "The bigger the interface, the weaker the abstraction."

3. Multiple small files vs. One large file: In Go, it is common to separate different things into different files. 
This can be different types, different interfaces, different functions, etc. It can be useful to separate interfaces 
into their own files named for what the interface represents. The package tree can give you a good idea of what 
interfaces are available without having to open a single huge file.

## Instructions
```shell
go mod init github.com/oskiegarcia/fetcher
```
### 2. Creat main.go

### 3. Run the program
```shell
go run .
```

