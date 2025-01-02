package main

import "fmt"

func creation() {
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
