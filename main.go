package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/structs_practices/note"
	"github.com/structs_practices/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Log()
}

func main() {

	todo, err := todo.New(getTodoData())

	if err != nil {
		fmt.Println("error when create todo: ", err)
		return
	}

	userNote, err := note.New(getNoteData())

	if err != nil {
		fmt.Println("error when create user note: ", err)
		return
	}

	fmt.Println("")

	err = saveData(todo)
	if err != nil {
		fmt.Println("error when save todo")
		return
	}

	fmt.Println("")

	err = saveData(userNote)
	if err != nil {
		fmt.Println("error when save note")
		return
	}
}

func saveData(data outputtable) error {
	data.Log()
	return data.Save()
}

func getNoteData() (string, string) {
	return getUserInput("Note Title: "), getUserInput("Note Content: ")
}

func getTodoData() string {
	return getUserInput("Todo Text: ")
}

func getUserInput(printMessage string) (input string) {
	fmt.Print(printMessage)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input

}
