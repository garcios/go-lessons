package main

import "fmt"

func passingFunc() {
	fmt.Println("function as argument type...")
	numbers := []int{1, 2, 3, 4}

	// pass the function definition as argument
	results := transformNumbers(&numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("double:", results)

	// define the function and store as variable
	triple := func(n int) int {
		return n * 3
	}

	// pass the variable function as argument
	results = transformNumbers(&numbers, triple)
	fmt.Println("triple:", results)
}

func transformNumbers(numbers *[]int, f transformFn) *[]int {
	nums := make([]int, len(*numbers))

	for i, num := range *numbers {
		nums[i] = f(num)
	}

	return &nums
}
