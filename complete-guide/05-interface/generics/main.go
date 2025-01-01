package main

import "fmt"

func main() {

	fmt.Println(add(1, 2)) //3

	fmt.Println(add(2.3, 2)) //4.3

	fmt.Println(add("Hello, ", "World")) //Hello, World

}

func add[T int | float64 | string](a T, b T) T {
	return a + b
}
