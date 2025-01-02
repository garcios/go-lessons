package main

import "fmt"

func growingList() {
	temp := []string{
		"Philippines",
		"Singapore",
		"Australia",
		"Germany",
		"Switzerland",
		"France",
		"Canada",
		"Vietnam",
		"Thailand",
		"Malaysia",
		"Indonesia",
		"China"}

	fmt.Printf("cap(temp):%d, addresss: %p\n", cap(temp), &temp[0]) // with capacity of 12
	temp = append(temp, "New Zealand")                              //adding 1 element will double the current capacity
	fmt.Printf("cap(temp):%d, addresss: %p\n", cap(temp), &temp[0])

	// my visited countries
	var countries []string

	for _, country := range temp {
		countries = append(countries, country)
		fmt.Printf("cap(countries):%d, addresss: %p\n", cap(countries), &countries[0])
	}

	/* Outputs - Note that a new array is allocated as capacity grows. It doubles as it grows.
	oscargarcia@Oscars-MBP-2 list % go run .
	cap(temp):12, addresss: 0x1400012a000
	cap(temp):24, addresss: 0x1400012e000
	cap(countries):1, addresss: 0x14000102020
	cap(countries):2, addresss: 0x14000130000
	cap(countries):4, addresss: 0x1400011c040
	cap(countries):4, addresss: 0x1400011c040
	cap(countries):8, addresss: 0x14000132000
	cap(countries):8, addresss: 0x14000132000
	cap(countries):8, addresss: 0x14000132000
	cap(countries):8, addresss: 0x14000132000
	cap(countries):16, addresss: 0x14000134000
	cap(countries):16, addresss: 0x14000134000
	cap(countries):16, addresss: 0x14000134000
	cap(countries):16, addresss: 0x14000134000
	cap(countries):16, addresss: 0x14000134000
	*/
}
