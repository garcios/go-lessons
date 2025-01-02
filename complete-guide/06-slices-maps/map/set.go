package main

import "fmt"

func set() {
	slice := []int{1, 2, 3, 4, 5, 5, 4, 3}
	set := sliceToSet(slice)

	for item := range set {
		fmt.Println(item)
	}

	value := 3
	if checkIfExists(set, value) {
		fmt.Printf("Element %d exists in the set\n", value)
	} else {
		fmt.Printf("Element %d does not exist in the set\n", value)
	}

	value = 10
	if checkIfExists(set, value) {
		fmt.Printf("Element %d exists in the set\n", value)
	} else {
		fmt.Printf("Element %d does not exist in the set\n", value)
	}
}

func sliceToSet(slice []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, item := range slice {
		set[item] = struct{}{}
	}

	return set
}

func checkIfExists(set map[int]struct{}, elem int) bool {
	_, exists := set[elem]
	return exists
}
