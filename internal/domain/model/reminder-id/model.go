package reminder_id_model

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type ReminderID int64

func (x ReminderID) Validate() error {
	if x == 0 {
		return errors.New("reminder id is zero")
	}

	return nil
}

func GenerateReminderID() ReminderID {
	return ReminderID(time.Now().Nanosecond())
}

func NewReminderID(n int64) (ReminderID, error) {
	x := ReminderID(n)
	if err := x.Validate(); err != nil {
		return 0, fmt.Errorf("failed to validate reminder id: %w", err)
	}

	return x, nil
}

func NewReminderIDFromString(s string) (ReminderID, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int: %w", err)
	}

	return NewReminderID(n)
}
