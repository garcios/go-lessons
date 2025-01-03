package main

import (
	"fmt"
	"strings"
)

func returnFunc() {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("function as return type...")
	numbers := []int{1, 2, 3, 4}

	fmt.Print("double: ")
	for _, num := range numbers {
		fmt.Print(getFunc("double")(num))
		fmt.Print(",")
	}
	fmt.Println()

	fmt.Print("triple: ")
	for _, num := range numbers {
		fmt.Print(getFunc("triple")(num))
		fmt.Print(",")
	}
	fmt.Println()

}

func getFunc(f string) transformFn {
	switch f {
	case "double":
		return func(n int) int {
			return n * 2
		}
	case "triple":
		return func(n int) int {
			return n * 3
		}
	default:
		return func(n int) int {
			return n
		}
	}

	return nil
}
