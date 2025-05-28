package reminder_title_model

import (
	"errors"
)

type ReminderTitle string

func NewReminderTitle(s string) (ReminderTitle, error) {
	x := ReminderTitle(s)
	if err := x.Validate(); err != nil {
		return "", err
	}

	return ReminderTitle(s), nil
}

func (x ReminderTitle) Validate() error {
	if x == "" {
		return errors.New("reminder_title is empty")
	}

	return nil
}
