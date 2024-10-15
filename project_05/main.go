package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saveToFiler interface {
	SaveToFile() error
}

type printer interface {
	Print()
}

type outputtable interface {
	saveToFiler
	printer
}

func main() {
	printSomething("1.5")
	printSomething(1)
	printSomething(5.5)

	noteTitle, noteContent := getUserData()
	text := getTodoData()


	userNote, noteError := note.New(noteTitle, noteContent)

	userTodo, todoError := todo.New(text)

	if noteError != nil {
		panic("Failed to create note!")
	}

	if todoError != nil {
		panic("Failed to create todo!")
	}

	userNote.Print()
	userTodo.Print()

	noteError = outputData(userNote)

	if noteError != nil {
		return 
	}

	todoError = outputData(userTodo)
}

// We can either use any or interface{}
func printSomething(data any) {
	intVal, ok := data.(int)

	if ok {
		fmt.Print(intVal)
		fmt.Println(" => Integer")
	}

	float64Val, ok := data.(float64)

	if ok {
		fmt.Print(float64Val)
		fmt.Println(" => Float64")
	}

	stringVal, ok := data.(string)

	if ok {
		fmt.Print(stringVal)
		fmt.Println(" => String")
	}

	// fmt.Print(data)
	//
	// switch data.(type) {
	// 	case string:
	// 		fmt.Println(" => String")
	// 	case float64:
	// 		fmt.Println(" => Float64")
	// 	case int:
	// 		fmt.Println(" => Integer")
	// 	default:
	// 		fmt.Println(" => Unknown")
	// }

}

func outputData(data outputtable) error {
	data.Print()
	return saveData(data)
}

func saveData(data saveToFiler) error {
	err := data.SaveToFile()

	if err != nil {
		return err
	}

	fmt.Println("Saved to file!")

	return nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	buffer := bufio.NewReader(os.Stdin)

	userInput, _ := buffer.ReadString('\n')

	userInput = strings.TrimSuffix(userInput, "\n")
	userInput = strings.TrimSuffix(userInput, "\r")

	return userInput
}

func getUserData() (string, string) {
	noteTitle := getUserInput("Note title: ")

	noteContent := getUserInput("Note content: ")

	return noteTitle, noteContent
}

func getTodoData() string {
	text := getUserInput("To-Do: ")

	return text
}
