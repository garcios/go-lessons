package main

import "fmt"

func main() {

	// Array has fixed number of elements
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

	// Slice - has dynamic number of elements
	products := []string{"macbook", "monitor", "headset"}
	fmt.Println("products:", products) // prints [macbook monitor headset]

	productList1 := products[1:2]              //slice
	fmt.Println("productList1:", productList1) // prints [monitor]

	products = append(products, "kindle")
	fmt.Println("products:", products) // prints [macbook monitor headset kindle]

	fmt.Println("len(products):", len(products)) //prints 4

	// loop through all products
	fmt.Println()
	for index, product := range products {
		fmt.Printf("%d-%s\n", index, product)
	}

	// Deep copy
	copiedProducts := make([]string, len(products))
	copy(copiedProducts, products)

	// modify the copy
	copiedProducts[0] = "macbook" + "2"
	fmt.Println("copiedProducts:", copiedProducts)
	fmt.Println("products:", products)

	// starting with nil slice
	var friends []string
	fmt.Println(friends == nil)                // prints true
	fmt.Println("friends:", friends)           // prints []
	fmt.Println("len(friends):", len(friends)) // prints 0
	fmt.Println("cap(friends):", cap(friends)) // prints 0

	friends = append(friends, "Doti", "Alan", "Rey") // creates the array when it does not exist
	fmt.Println("friends:", friends)                 // prints [Doti Alan Rey]
	fmt.Println("len(friends):", len(friends))       // prints 3
	fmt.Println("cap(friends):", cap(friends))       // prints 3

	// using make
	cities := make([]string, 0, 5)           // initial capacity of 5
	fmt.Println(cities == nil)               // prints false
	fmt.Println("cities:", cities)           // prints []
	fmt.Println("len(cities):", len(cities)) // prints 0
	fmt.Println("cap(cities):", cap(cities)) // prints 5

	cities = append(cities, "Melbourne", "Sydney", "Brisbane")
	fmt.Println("cities:", cities)           // prints [Melbourne Sydney Brisbane]
	fmt.Println("len(cities):", len(cities)) // prints 3
	fmt.Println("cap(cities):", cap(cities)) // prints 5

	moreCities := []string{"Manila", "Makati", "Pasig"}
	cities = append(cities, moreCities...)   //adding 3 more will reach capacity and therefore will create new array
	fmt.Println("cities:", cities)           // prints [Melbourne Sydney Brisbane Manila Makati Pasig]
	fmt.Println("len(cities):", len(cities)) // prints 6
	fmt.Println("cap(cities):", cap(cities)) // prints 10

}
