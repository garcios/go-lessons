package main

import (
	"fmt"
	"strings"
)

func iterate() {
	languages := []string{"Go", "Java", "Python", "Javascript"}

	// using the value
	for _, language := range languages {
		fmt.Println(language)
	}

	fmt.Println(strings.Repeat("-", 50))

	// using the index
	for i := range languages {
		fmt.Println(languages[i])
	}
}
