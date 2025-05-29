package reminder_id_model

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type ReminderID int

func (x ReminderID) Validate() error {
	if x == 0 {
		return errors.New("reminder id cannot be empty")
	}

	return nil
}

func GenerateReminderID() ReminderID {
	return ReminderID(time.Now().Nanosecond())
}

func NewReminderID(s string) (ReminderID, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int: %w", err)
	}

	x := ReminderID(n)
	if err = x.Validate(); err != nil {
		return 0, fmt.Errorf("failed to validate reminder: %w", err)
	}

	return x, nil
}
