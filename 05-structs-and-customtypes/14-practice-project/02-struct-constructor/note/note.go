package note

import (
	"errors"
	"time"
)

type Note struct {
	title    string
	content  string
	createAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		title:    title,
		content:  content,
		createAt: time.Now(),
	}, nil
}
