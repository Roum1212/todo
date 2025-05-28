package reminder_title_model

import (
	"errors"
	"fmt"
)

type ReminderTitle string

func (x ReminderTitle) Validate() error {
	if x == "" {
		return errors.New("reminder title is empty")
	}

	return nil
}

func NewReminderTitle(s string) (ReminderTitle, error) {
	x := ReminderTitle(s)
	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("failed to validate reminder title: %w", err)
	}

	return x, nil
}

func MustNewReminderTitle(s string) ReminderTitle {
	return ReminderTitle(s)
}
