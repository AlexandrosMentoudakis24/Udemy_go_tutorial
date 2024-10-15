package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		fmt.Println("Text is required!")

		return nil, errors.New("Failed to create todo!")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (todo *Todo) Print() {
	fmt.Println("To-Do:", todo.Text)
}

func (todo *Todo) SaveToFile() error {
	fmt.Println("Saving to file!")

	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
