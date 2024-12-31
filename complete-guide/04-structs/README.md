# Structs
In Go programming language, a struct (short for structure) is a composite data type that groups together zero or more 
values of different types. These grouped values are called fields. Each field has a unique name and a specific type

Structs are very useful when you need to represent entities or objects that have multiple properties. They allow you to 
group related data together, making your code cleaner and more organized.

Here is a simple example of a struct in Go:
```go
type Book struct {
    Title  string
    Author string
    Pages  int
}
```
In this example, Book is a struct type that has three fields: Title and Author (which are of type string) and Pages 
(which is of type int).

You can create an instance of a struct like this:
```go
var myBook Book
myBook.Title = "Go Programming"
myBook.Author = "Example Author"
myBook.Pages = 320
```
Or you can create and initialize a struct in the same line:
```go
myBook := Book{
    Title:  "Go Programming",
    Author: "Example Author",
    Pages:  320,
}
```
Then, you can access the fields of the struct using dot notation:
```go
fmt.Println(myBook.Title)  // prints: Go Programming
```
Structs can also contain methods. A method is a function that is associated with a particular struct type. It's declared
with a receiver, which specifies the struct type that the method is associated with. For example:

```go
func (b Book) describe() {
    fmt.Printf("%s by %s has %d pages.\n", b.Title, b.Author, b.Pages)
}
```
You can call this method on instances of the Book struct:
```go
myBook.describe()  // prints: Go Programming by Example Author has 320 pages.
```

## Encapsulation
In Go programming language, encapsulation is achieved using packages and the visibility rules they provide. You can 
designate fields as unexported (lowercased initial letter) in your structs to prevent them from being directly accessed
from outside the package they're declared in. This encapsulation is a core part of good object-oriented design.

However, remember this is not exactly "hiding" in the pure Object Oriented Programming (OOP) sense, as Go does not 
support classes or private/public modifiers in the way some other languages like Java and C++ do.

Here is an example:

```go
package mypackage

type myStruct struct {
    MyExportedField   int    // This field is visible outside 'mypackage'
    myUnexportedField string // This field is not visible outside 'mypackage'
}
```

With this, you can control which parts of your structs are accessible from other packages, and which ones are not. 
It's not exactly "using a package to hide struct variables," but rather utilizing how Go's package system works to 
attain desired encapsulation.

Organizing your code into packages can be a very good idea - it can lead to more maintainable and understandable code.

It's good to keep in mind that to provide access to unexported fields, getter and setter methods are commonly used. 
They allow you to control how an object's variables are accessed and can provide additional logic around setting and 
getting the values.

Example:
```go
package mypackage

type myStruct struct {
    myUnexportedField string // This field is not visible outside my package
}

// Getter
func (ms *myStruct) GetField() string {
    return ms.myUnexportedField
}

// Setter
func (ms *myStruct) SetField(val string) {
    // additional logic can go here
    ms.myUnexportedField = val
}
```

## Embedding
In Go, embedding is a way to reuse code by including an existing type into another type. It allows us to borrow the 
fields and methods from one struct type into another. However, unlike inheritance in many object-oriented languages, 
embedding in Go is more about composition rather than an "is-a" relationship.

Let's look at an example of embedding:
```go
type Person struct {
    Name string
    Age  int
}

// The struct 'Student' embeds struct 'Person'
type Student struct {
    Person
    MatriculationNumber int
}

// Now we can use the properties and methods of 'Person' on 'Student'
s := Student{
    Person: Person{
        Name: "John Doe",
        Age: 21,
    },
    MatriculationNumber: 12345678,
}

fmt.Println(s.Name)   // prints: John Doe
fmt.Println(s.Age)    // prints: 21
fmt.Println(s.MatriculationNumber)   // prints: 12345678
```

In this example, Person is an embedded field in the Student struct. So, Student now has access to all the fields and 
methods that Person had.

When embedding, the methods of the outer type take precedence. Thus if the outer type and the embedded type have methods
with the same name, calls to that method name will invoke the outer type's method.

Another important thing to note is that Go does not support classic OOP; it has no concept of classes or inheritance. 
Instead, Go utilizes structs, interfaces, and embedding to achieve flexible and powerful polymorphic behavior.

## Alias
In Go, an alias type is a way to create a new name for an existing type. The alias type has the same underlying 
structure as the original type. This can be useful for improving the readability of the code or for creating 
semantically meaningful types.

Here's an example of creating type alias:
```go
type Celsius float64
type Fahrenheit float64
```

In this example, both Celsius and Fahrenheit are float64 types, but they have been aliased to give them more context in 
a program that needs to distinguish between temperatures in Celsius and Fahrenheit.

You can create functions that accept and return these types:
```go
func (c Celsius) toF() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) toC() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

func main(){
	var c Celsius
	var f Fahrenheit
	
	c = 100
	
	f = c.toF()
}

```

Note that even though Celsius and Fahrenheit are both underlined by float64, they are not interchangeable because they 
are different types. So the following code will trigger a compiler error:
```go
var c Celsius = 0
var f Fahrenheit
f = c // Compiler error: cannot use c (type Celsius) as type Fahrenheit in assignment
```
The alias type only lasts for the scope in which it was declared, so it only works throughout the package if it was 
declared in the package scope.