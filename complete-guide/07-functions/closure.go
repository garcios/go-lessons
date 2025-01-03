package main

import "fmt"

func closure1() {
	counter := count()
	fmt.Println(counter()) // prints 1
	fmt.Println(counter()) // prints 2
	fmt.Println(counter()) // prints 3
}

func count() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Counter struct {
	Value int
}

// Function to produce a closure
func CreateCounter() func() int {
	counter := &Counter{Value: 100}
	return func() int {
		counter.Value++
		return counter.Value
	}
}

func closure2() {
	// Get the function, sets initial valur to 100
	counter1 := CreateCounter()

	// The closure has access to the Counter struct's Value
	fmt.Println(counter1()) // Prints 101
	fmt.Println(counter1()) // Prints 102
	fmt.Println(counter1()) // Prints 103

	// Get the function, it will have a separate state, i.e. value is 100 again.
	counter2 := CreateCounter()

	// The closure has access to the Counter struct's Value
	fmt.Println(counter2()) // Prints 101
	fmt.Println(counter2()) // Prints 102
	fmt.Println(counter2()) // Prints 103
}

func closure3() {
	x := 10
	foo := func() {
		// The closure captures the current value of x
		fmt.Println(x)
		x++
	}
	x = 20
	foo() // This will print 20, not 10
	foo() // This will print 21
}

// the return function has access to variable factor.
func createTransformer(factor int) transformFn {
	return func(num int) int {
		return num * factor
	}
}

func closure4() {
	nums := []int{1, 2, 3, 4, 5}
	double := createTransformer(2)
	triple := createTransformer(3)

	for _, num := range nums {
		fmt.Println(double(num))
	}

	fmt.Println()

	for _, num := range nums {
		fmt.Println(triple(num))
	}

}
