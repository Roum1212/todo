package reminder_id_model

import (
	"time"
)

type ReminderID int

func GenerateReminderID() ReminderID {
	return ReminderID(time.Now().Nanosecond())
}
