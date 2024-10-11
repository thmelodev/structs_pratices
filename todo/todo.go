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
		return &Todo{}, errors.New("Todo text should not be empty")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (todo *Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	os.WriteFile(filename, json, 0644)

	return nil
}

func (todo *Todo) Log() {
	fmt.Printf("Todo: %v \n", todo.Text)
}
