package main

import "fmt"

func array() {
	prices := [5]float32{100, 200, 300, 400, 500}
	fmt.Println("prices:", prices) // prints [100 200 300 400 500]

	// Note: first index is inclusive, and second indexed is exclusive.
	priceList1 := prices[2:4]              //slice
	fmt.Println("priceList1:", priceList1) // prints [300 400]

	priceList2 := prices[:4]               //slice
	fmt.Println("priceList2:", priceList2) // prints [100 200 300 400]

	priceList3 := prices[2:]               //slice
	fmt.Println("priceList3:", priceList3) // prints [300 400 500]

	priceList3[0] = 301            // this updates the underlying array - prices
	fmt.Println("prices:", prices) // prints [100 200 301 400 500]

	// updates the 5th element
	prices[4] = 501
	fmt.Println("prices:", prices) // prints [100 200 300 400 501]

	fmt.Println("prices[0]:", prices[0]) // prints 100
}
