package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note uses struct tags for the json output.
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// New is constructor function.
func New(title string, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	if content == "" {
		return nil, errors.New("content must not be empty")
	}

	return &Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}

func (n *Note) Display() {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Title: %s\n", n.Title)
	fmt.Printf("Content: %s\n", n.Content)
	fmt.Printf("CreatedAt: %s\n", n.CreatedAt.Format(time.RFC3339))
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
