package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/structs_practices/note"
)

func main() {
	title, content := getUserInput("Note Title: "), getUserInput("Note Content: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("error when create user note: ", err)
		return
	}

	userNote.Save()
	userNote.Log()
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
