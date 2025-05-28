package reminder_description_model

import (
	"errors"
)

type ReminderDescription string

func NewReminderDescription(s string) (ReminderDescription, error) {
	x := ReminderDescription(s)
	if err := x.Validate(); err != nil {
		return "", err
	}

	return ReminderDescription(s), nil
}

func (x ReminderDescription) Validate() error {
	if x == "" {
		return errors.New("reminder_description is empty")
	}

	return nil
}
