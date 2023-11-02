package day

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (event *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	event.title = title
	return nil
}
