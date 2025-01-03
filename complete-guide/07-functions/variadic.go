package main

import "fmt"

func variadic() {
	fmt.Println(sum(1, 2))          // prints 3
	fmt.Println(sum(1, 2, 3))       // prints 6
	fmt.Println(sum(1, 2, 3, 4, 5)) //prints 15

	nums := []int{1, 2, 3, 4, 5}

	// splitting slice into parameter values using ... operator
	fmt.Println(sum(nums...)) //prints 15

	fmt.Println("Hello", "Oscar", "How", "are", "you?") //prints "Hello Oscar How are you?"

}

// nums is a slice
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
