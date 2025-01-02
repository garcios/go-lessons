package main

import "fmt"

func main() {

	// create using map literal that is empty
	//creation1()

	// create using map literal that is initialized with some values
	//creation2()

	// create using make
	//creation3()

	// with struct as valus
	creation4()

	// accessing map
	//access()

	// delete from map
	//deleteElement()

	// set()
}

func creation1() {
	websites := map[string]string{}
	websites["Google"] = "www.google.com"
	websites["Amazon"] = "www.amazon.com"

	for k, v := range websites {
		fmt.Println(k, v)
	}

	fmt.Println(websites["meta"]) // prints www.facebook.com
}

func creation2() {
	students := map[int]string{
		1: "Oscar",
		2: "Pamela",
	}

	for k, v := range students {
		fmt.Println(k, v)
	}
}

func creation3() {
	stockPrices := make(map[string]float64, 3) // optionally, we can specify the capacity
	stockPrices["AMZN"] = 219.39
	stockPrices["GOOGL"] = 189.30
	stockPrices["MSFT"] = 421.50

	for k, v := range stockPrices {
		fmt.Println(k, v)
	}
}

func creation4() {
	type product struct {
		id    string
		name  string
		price float64
	}

	// using an alias for a map
	type productMap map[string]*product

	newProduct := func(id string, name string, price float64) *product {
		return &product{id: id, name: name, price: price}
	}

	products := productMap{
		"123": newProduct("123", "Laptop", 999.99),
		"234": newProduct("234", "Monitor", 500.00),
	}

	for k, v := range products {
		fmt.Println(k, v)
	}

}

func access() {
	fruitPrices := map[string]float64{
		"apple":  1.0,
		"orange": 2.0,
		"grape":  3.0,
	}

	_, ok := fruitPrices["avocado"]
	if !ok {
		fmt.Println("no avocado")
	}

}

func deleteElement() {
	fruitPrices := map[string]float64{
		"apple":  1.0,
		"orange": 2.0,
		"grape":  3.0,
	}

	delete(fruitPrices, "apple")

	fmt.Println(fruitPrices)
}
