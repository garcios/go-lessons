package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Todo uses struct tags for the json output.
type Todo struct {
	Text string `json:"text"`
}

// New is constructor function.
func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("text must not be empty")
	}

	return &Todo{Text: text}, nil
}

func (t *Todo) Display() {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Text: %s\n", t.Text)
}

func (t *Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
