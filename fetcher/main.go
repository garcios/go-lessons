package main

import (
	"fmt"
)

// Fetcher defines the two methods we need to fetch members for each house.
type Fetcher interface {
	GetAll() ([]string, error)
	GetSourceName() string
}

// verify interface compliance
var _ Fetcher = (*Gryffindor)(nil)
var _ Fetcher = (*Hufflepuff)(nil)
var _ Fetcher = (*Ravenclaw)(nil)
var _ Fetcher = (*Slytherin)(nil)

func main() {

	err := fetchAndDisplayAllMembers()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

}

func fetchAndDisplayAllMembers() error {
	fetchers := []Fetcher{
		NewGryffindor(),
		NewSlytherin(),
		NewRavenclaw(),
		NewHufflepuff(),
	}

	for _, fetcher := range fetchers {
		results, err := fetcher.GetAll()
		if err != nil {
			return fmt.Errorf("GetAll error: %v\n", err)
		}

		for _, result := range results {
			fmt.Printf("%s, Fetched member: %s\n", fetcher.GetSourceName(), result)
		}
	}

	return nil
}
