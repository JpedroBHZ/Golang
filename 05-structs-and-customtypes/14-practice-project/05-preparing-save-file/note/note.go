package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title    string
	content  string
	createAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note tilted %v has the following content: \n\n%v\n\n", note.title, note.content)
}

func (note note ) Save {
	fileName := strings.ReplaceAll(note.title, "", "_")
	fileName = strings.ToLower(fileName)
	os.WriteFile(fileName, )
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title:    title,
		content:  content,
		createAt: time.Now(),
	}, nil
}
