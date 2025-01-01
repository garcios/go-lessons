# Map
Maps are a key-value data structure available in Go. They can store items of various data types and allow you to quickly
access the stored values using the corresponding keys. Similar to dictionaries in Python or objects in JavaScript, the 
keys in a map are unique, and each key maps to exactly one value.

Here is the basic syntax to declare a map:
```go
var mapName map[keyType]valueType
```

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

