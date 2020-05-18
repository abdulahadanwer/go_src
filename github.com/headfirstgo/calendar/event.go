package calendar

import (
	"errors"
	"unicode/utf8"
)

// Event holds event title with a date
type Event struct {
	title string
	Date
}

// Title returns the title set
func (e *Event) Title() string {
	return e.title
}

// SetTitle sets the title of an event
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid Title")
	}
	e.title = title

	return nil
}
