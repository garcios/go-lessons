package main

import "fmt"

func slice() {
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

	// Removing first element from slice
	prices1 := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println("prices1:", prices1)

	prices1 = prices1[1:]
	fmt.Println("prices, 1st element removed:", prices1)

	// Removing last element from slice
	prices2 := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println("prices2:", prices2)

	prices2 = prices2[:len(prices2)-1]
	fmt.Println("prices2, last element removed:", prices2)

	// Removing 3rd element from slice
	prices3 := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println("prices3:", prices3)

	temp := prices3[:2]
	temp = append(temp, prices3[3:]...) //merging a slice by unpacking it with ... operator
	prices3 = temp
	fmt.Println("prices3, 3rd element removed:", prices3)
}
