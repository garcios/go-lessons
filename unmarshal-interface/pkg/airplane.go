package pkg

import (
	"fmt"
)

type Airplane struct {
	Type     string `json:"type"`
	Model    string `json:"model"`
	Capacity int    `json:"capacity"`
}

func (a *Airplane) Info() {
	fmt.Println(a)
}
