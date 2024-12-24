package main

import (
	"fmt"
)

type Fetcher interface {
	GetByID(int) (string, error)
	GetSourceName() string
}

// verify interface compliance
var _ Fetcher = (*Gryffindor)(nil)
var _ Fetcher = (*Hufflepuff)(nil)
var _ Fetcher = (*Ravenclaw)(nil)
var _ Fetcher = (*Slytherin)(nil)

func main() {

	fetchers := []Fetcher{
		&Gryffindor{Data: []string{"Harry", "Ron", "Hermione"}},
		&Slytherin{Data: []string{"Draco", "Crabbe", "Goyle"}},
		&Ravenclaw{Data: []string{"Luna", "Cho", "Padma"}},
		&Hufflepuff{Data: []string{"Cedric", "Nymphadora", "Teddy"}},
	}

	// assume we know the IDs
	IDs := []int{0, 1, 2}

	for _, fetcher := range fetchers {
		for _, id := range IDs {
			ret, err := fetcher.GetByID(id)
			if err != nil {
				fmt.Printf("Error fetching by id %d: %v\n", id, err)
				return
			}
			fmt.Printf("%s, Fetched with id %d: %s\n", fetcher.GetSourceName(), id, ret)
		}
	}

}
