package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Record struct {
	ID      int
	Name    string
	Created time.Time
}

func main() {
	names := make([]Record, 0)

	for i := 1; i <= 17; i++ {
		names = append(names, createRecord(i, fmt.Sprintf("name_%d", i)))
	}

	// list record
	dateFormat := "2006-01-02 15:04:05"
	pageNumbers := []int{1, 2, 3, 4}

	sort.Slice(names, func(i, j int) bool {
		return names[i].Created.After(names[j].Created)
	})

	for _, n := range pageNumbers {
		fmt.Printf("[Page %d]------------------------\n", n)
		page := paginate(names, n, 5)
		for _, p := range page {
			fmt.Printf("%v: %v\n", p.Name, p.Created.Format(dateFormat))
		}
	}

}

func createRecord(id int, name string) Record {
	now := time.Now()
	daysToAdd := getRandomNumber()

	return Record{
		ID:      id,
		Name:    name,
		Created: now.AddDate(0, 0, daysToAdd),
	}
}

func paginate(records []Record, currentPage int, pageSize int) []Record {
	start := (currentPage - 1) * pageSize
	end := start + pageSize

	if start >= len(records) {
		return []Record{}
	}

	if end > len(records) {
		end = len(records)
	}

	return records[start:end]
}

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
