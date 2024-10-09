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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {

	if title == "" {
		return nil, errors.New("note title should not be empty")
	}
	if content == "" {
		return nil, errors.New("note content should not be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	os.WriteFile(filename, json, 0644)

	return nil
}

func (note *Note) Log() {
	fmt.Println("\n------------ Your Note ------------\n")
	fmt.Printf("Note title: %v \n", note.Title)
	fmt.Printf("Note Content: %v \n", note.Content)
	fmt.Printf("Note CreatedAt: %v \n", note.CreatedAt.Format("02/01/2006 15:04:05"))
}
