# Array
An array in Go is a fixed-length sequence of items of the same type. The items are arranged in a continuous block of 
memory, and you can access each item using an integer index, which starts from 0.

In Go, arrays are defined using the following syntax:
```go
var arrayName [size]Type
```
Where arrayName is the name of the array, size is the number of items in the array, and Type is the data type of the 
items.

Here's an example of declaring, initializing, and iterating over an array in Go:
```go
package main

import "fmt"

func main() {
    var numbers [5]int // Declare an array of 5 integers
    numbers[0] = 1     // Assign a value to the first element
    numbers[1] = 2     // Assign a value to the second element
    // The rest of the elements will have a zero value

    for i, number := range numbers {
        fmt.Printf("numbers[%d] = %d\n", i, number)
    }
}
```
A critical aspect of Go's arrays is that their size is part of their type. That means arrays of different sizes are 
considered to be different types.

## Slice
A slice in Go is a flexible, growable, and more commonly used alternative to arrays. It doesn't have a fixed size, 
allowing elements to be dynamically added or removed. A slice is a segment of an array, and it gives more power, 
flexibility, and convenience compared to arrays. While arrays are of fixed size, slices are dynamic.

Underneath, a slice is just a reference to an underlying array. As such, a slice consists of three main components:
1. A pointer to the underlying array.
2. The length of the slice.
3. The capacity of the slice (the maximum length the slice can reach before new memory has to be allocated).

The syntax to declare a slice in Go is:
```go
var sliceName []Type
```
Where sliceName is the name of the slice, and Type is the data type of the items.

Here's an example of using a slice in Go:
```go
package main

import "fmt"

func main() {
    // Declare a slice (this slice is nil by default)
    var fruits []string

    // Append some values to the slice
    fruits = append(fruits, "apple", "banana", "cherry")

    for i, fruit := range fruits {
        fmt.Printf("fruits[%d] = %s\n", i, fruit)
    }
}
```
In this example, fruits is a slice that holds strings. We use the append function to add elements to the slice. 
Unlike arrays, slices grow dynamically as we add more elements.

## Underlying array of a slice
__Question__: Given the following code, does priceList1 create a new underlying array?

```go
// Array has fixed number of elements
prices := [5]float32{100, 200, 300, 400, 500}
fmt.Println(prices) // prints [100 200 300 400 500]

priceList1 := prices[2:4] //slice
fmt.Println(priceList1)   
```

__Answer__:
No, in the provided code priceList1 doesn't create a new underlying array. Instead, it creates a slice from the prices 
array. The slice references the same underlying array and simply points to a different range within that array.

In the above code, priceList1 is a slice that includes the elements at index 2 and 3 of the prices array. The resulting 
slice includes the values 300 and 400, but the underlying array is still the original prices array.

Remember, slices are just views into an array. So, they don't contain any data themselves, they just describe a section
of the underlying array. When you change an element in a slice, you're actually changing the element in the underlying 
array. Similarly, if you change an element in the underlying array, that change will be visible in the slices.

## Deep copy
To create a deep copy of a subarray (or a slice) in Go, you could use the copy function. The copy function copies 
elements from a source slice to a destination slice and returns the number of elements copied.

First, you would need to make a new slice with the necessary length, and then use the copy function to copy the 
elements. Here is an example:

```go
package main

import "fmt"

func main() {
    // Original array
    prices := [5]float32{100, 200, 300, 400, 500}

    // Subarray (slice)
    priceList1 := prices[2:4]

    // Make a new slice with the necessary length
    priceListCopy := make([]float32, len(priceList1))

    // Use the copy function to perform a deep copy
    copy(priceListCopy, priceList1)
	priceListCopy[0]=301 // modify to prove that it will not affect the original array.
	
    fmt.Println(priceList1)   // prints [300 400]
    fmt.Println(priceListCopy) // prints [301 400]
}
```

## var mySlice []string vs make([]string, 0)
There is a subtle difference between declaring a slice in Go using []string{} or var mySlice []string and 
using mySlice := make([]string, 0).

1. var mySlice []string or mySlice := []string{} declares a nil slice. A nil slice does not have an underlying array, 
its length and capacity are 0, and it has no memory allocated. In other words this slice points to no allocated memory.
However, you can use append() to add elements to it because append() checks if the slice is nil and accordingly 
allocates memory.

2. mySlice := make([]string, 0) declares an empty slice but not a nil slice. This means the slice has an underlying 
array but no elements, its length is 0, but its capacity can be greater than 0 if specified in the make function 
(like in mySlice := make([]string, 0, 5)). make() allocates memory for the slice, so it's not pointing to nil but an 
empty segment of memory.

These two different ways to declare a slice can behave identically when used (especially with functions like append()), 
but they can show different behaviors in certain scenarios - such as when they are compared directly to nil.

Example:
```go
var s1 []string
s2 := []string{}
s3 := make([]string, 0)

fmt.Println(s1 == nil)  // true
fmt.Println(s2 == nil)  // false
fmt.Println(s3 == nil)  // false
```
As shown in the example, only s1 is equal to nil. Even though s2 and s3 are empty, since they have memory allocated 
(they are not nil), they are not equal to nil.
