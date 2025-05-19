package reminder_id_model

import (
	"fmt"
	"strconv"
	"time"
)

type ReminderID int

func GenerateReminderID() ReminderID {
	return ReminderID(time.Now().Nanosecond())
}

func NewReminderID(s string) (ReminderID, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("reminder_id cannot be parsed as int: %w", err)
	}

	return ReminderID(n), nil
}
