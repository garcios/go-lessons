# Interface
In Go, an interface is a type definition that specifies a method set (collection of methods), without the actual 
implementation of those methods. An interface can be implicitly implemented by any type that has those methods.

Unlike some other languages, you don't explicitly declare that a certain type satisfies an interface in Go – merely 
possessing the necessary methods is enough.

## Conventions
- If an interface has only one method, then the name is the method name with "r".

For example:
```go
type saver interface{
	Save() error
}
```

## Best practices
Interfaces in Go are a powerful tool, but their usage requires thought to prevent unnecessary complexity in your code. 
Following are some best practices that are often advisable when working with Go interfaces:

1. __Use interfaces to express behavior__: Go interfaces are best used to express behaviors of objects, through methods. 
Therefore, design your interfaces around what things do, not what they are.

2. __Keep interfaces small__: The smaller an interface, the easier it is to use and implement. The Go standard library,
for instance, has interfaces with only one or two methods (like io.Reader, io.Writer). This practice known as "interface
segregation principle" in SOLID principles.

3. __Don't export interfaces for types__: If you have defined a type that is going to be used by other packages, don't 
define an interface for it to create a contract. The calling code should define the interface according to what methods 
it requires.

4. __Use very few or no pointer receiver interfaces__: Methods on interfaces should be designed so as not to modify the 
state of the object; this enables more flexible usage of your interfaces.

5. __Avoid using interfaces for sake of generality or future-proofing__: Don't use interfaces to make your code 
"future proof". Avoid making assumptions about the future requirements of your code.

Remember, these guidelines are a starting point, and like all rules, there can be exceptions depending on your specific
problem space.

## Embedded Interfaces
In Go, interfaces themselves can include (or embed) other interfaces. This means that one interface can include the 
method set of another, which is a powerful way to compose interfaces and build upon existing abstractions.

When we say an interface is embeddable, we mean that the interface can be included in other interfaces or structs. 
This inclusion is made without giving a specific field name.

Here is an example:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter interface embeds Reader and Writer
type ReadWriter interface {
    Reader  // it indicates that ReadWriter includes methods of Reader
    Writer  // it indicates that ReadWriter includes methods of Writer
}
```