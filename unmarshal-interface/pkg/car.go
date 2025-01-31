package pkg

import (
	"fmt"
)

type Car struct {
	Type  string `json:"type"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (c *Car) Info() {
	fmt.Println(c)
}
