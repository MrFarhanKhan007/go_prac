package noteData

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const path = "FileHandlingFiles\\"

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("\nTitle: %v\nContent: %v\nCreated At: %v", note.Title, note.Content, note.CreatedAt)
}

func (note Note) Save() (err error) {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		fmt.Printf("Error while converting data to JSON!\n%v", err)
		return
	}
	FullfileName := path + fileName
	return os.WriteFile(FullfileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid Data")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
