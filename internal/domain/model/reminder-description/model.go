package reminder_description_model

import (
	"errors"
	"fmt"
)

type ReminderDescription string

func (x ReminderDescription) Validate() error {
	if x == "" {
		return errors.New("reminder description is empty")
	}

	return nil
}

func NewReminderDescription(s string) (ReminderDescription, error) {
	x := ReminderDescription(s)
	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("failed to validate reminder title: %w", err)
	}

	return x, nil
}
