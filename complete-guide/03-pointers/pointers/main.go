package main

import "fmt"

func main() {

	age := 32
	var agePointer *int

	fmt.Println("age:", age)               // 32
	fmt.Println("agePointer:", agePointer) // nil

	// assign address of age variable
	agePointer = &age

	fmt.Println("agePointer:", agePointer) // 0x1400000e098

	// dereference
	fmt.Println("*agePointer:", *agePointer) // 32

	// passing pointer to a function
	adultYears := getAdultYears(agePointer)
	fmt.Println("adultYears:", adultYears)

	mutateAge(agePointer)
	fmt.Println("age after mutation:", age)
}

// pointer declaration in a function argument.
func getAdultYears(age *int) int {
	return *age - 18
}

func mutateAge(age *int) {
	*age = *age - 18
}
