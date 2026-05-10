package noteApp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrFarhanKhan007/go_prac/noteApp/noteData"
)

func NoteApp() {
	// get title, content from userInput with some pretext
	// store the resul in .json with using structs dataType as well

	title, content := getNoteData()

	userNote, err := noteData.New(title, content)

	if err != nil {
		fmt.Printf("Creating the note failed!\n%v", err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err!=nil {
		fmt.Printf("\nSaving the note Failed!\n%v", err)
		return
	}
	fmt.Println("\nSaving the note succeeded!")
}

func getNoteData() (title string, content string) {
	title = getUserValue("Enter Note Title: ")
	content = getUserValue("Enter Note Content: ")

	return title, content
}

func getUserValue(text string) (value string) {
	fmt.Print(text)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err!=nil {
		fmt.Printf("Error while reading input!\n%v", err)
		return
	}

	value = strings.TrimSuffix(value,"\n")
	value = strings.TrimSuffix(value,"\r")

	return value
}
