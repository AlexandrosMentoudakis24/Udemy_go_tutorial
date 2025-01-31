package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		fmt.Println("Title and content are required!")

		return nil, errors.New("Failed to create note!")
	}

	return &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Print() {
	fmt.Println("Title:", note.Title)
	fmt.Println("Content:", note.Content)
	fmt.Println("Created at:", note.CreatedAt)
}

func (note *Note) SaveToFile() error {
	fmt.Println("Saving to file!")

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
