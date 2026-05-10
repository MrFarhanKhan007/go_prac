package noteData

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note *Note) Display() {
	fmt.Printf("\nTitle: %v\nContent: %v\nCreated At: %v", note.title, note.content, note.createdAt)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid Data")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
