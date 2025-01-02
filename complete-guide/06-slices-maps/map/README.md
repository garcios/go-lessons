# Map
Maps are a key-value data structure available in Go. They can store items of various data types and allow you to quickly
access the stored values using the corresponding keys. Similar to dictionaries in Python or objects in JavaScript, the 
keys in a map are unique, and each key maps to exactly one value.

Here is the basic syntax to declare a map:
```go
var mapName map[keyType]valueType
```

## Creation
In Go, maps can be created in two ways - using make and using a map literal.
Here are the differences between the two:

1. Using make:
```go
m := make(map[string]string)
```
- The make function creates a map with no entries.
- It allocates and initializes a hash map data structure and returns a map value that points to it.
- The map is ready to accept key-entry values.
- It can also take a second, optional argument: a hint to the number of entries the map will hold, which helps in 
optimizing the memory allocation.
```go
m := make(map[string]string, 10)
```

2. Using a map literal:
```go
m := map[string]string{}
```
- A map literal allows a map to be created and initialized in the same line.
- A map literal has an underlying array and creates an entry for each key-entry pair. In the above case, there are no 
entries as it is an empty map literal.
- A map literal can also be initialized with key-entry values as such:
```go
m := map[string]string{"Hello": "world", "foo": "bar"}
```
So, both make(map[string]string) and map[string]string{} would create an empty map, but map literals are more versatile
as they can be used to initialize the map with key-entry pairs.

In performance, both are practically the same when creating an empty map. But when initializing with data, using a map 
literal can be more performant, because it creates an array of the right size from the start. If you use make and then 
insert key-value pairs one by one into a map, the map may have to be resized and rehashed multiple times, resulting in a performance hit.

## Set
In Golang, there is no built-in set data type, but we can use a map to simulate a set. 
The keys of the map can represent the elements of the set, and you can use empty struct (struct{}) as map values since 
they don't occupy any memory. Here is how you can convert a slice to a set:

```go
package main

import "fmt"

func sliceToSet(slice []int) map[int]struct{} {
    set := make(map[int]struct{})
    for _, item := range slice {
        set[item] = struct{}{}
    }
    return set
}

func main() {
    slice := []int{1, 2, 3, 4, 5, 5, 4, 3}
    set := sliceToSet(slice)

    for item := range set {
        fmt.Println(item)
    }
}
```

In this example, sliceToSet() function accepts a slice and creates a map. It then loops over the slice, using the slice 
values as keys in the map. The values in the map are of the empty struct type. In the main function, you can then print 
out the items in the resulting "set", and you'll see that duplicate items from the slice do not appear in the map, 
effectively creating a set.

