package main

import (
	"bufio"
	"fmt"
	"github.com/oskiegarcia/todo-app/note"
	"github.com/oskiegarcia/todo-app/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

// embed interface
type outputtable interface {
	saver
	Display()
}

func main() {

	text := getTodoData()

	userTodo, err := todo.New(text)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}

	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}

	fmt.Println("Saving succeeded!")

	return nil
}

func getTodoData() string {
	text := getUserInput("Todo text: ")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}
