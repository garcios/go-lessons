package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

func main() {
	print(1)              // prints int: 1
	print(1.5)            // prints float64: 1.5
	print("hello")        // prints string: hello
	print(true)           //prints unsupported: bool
	print(Person{"John"}) //prints unsupported: main.Person

	cast(3)   //prints Successfully casted to int: 3
	cast(3.1) //prints Unsupported type: float64
}

func print(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Printf("unsupported: %T\n", v)
	}
}

// any is an alias of interface{}
func cast(value any) {
	val, ok := value.(int) //cast value to int type
	if ok {
		fmt.Println("Successfully casted to int:", val)
	} else {
		fmt.Println("Unsupported type:", reflect.TypeOf(value))
	}
}
