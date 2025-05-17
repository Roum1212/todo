package reminder_title_model

import (
	"errors"
	"fmt"
)

type ReminderTitle string

func (x ReminderTitle) Validate() error {
	if len(x) > 5 {
		return errors.New("reminder title is too long")
	}

	return nil
}

func NewReminderTitle(s string) (ReminderTitle, error) {
	x := ReminderTitle(s)

	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("invalid reminder title: %w", err)
	}

	return x, nil
}
