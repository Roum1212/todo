package reminder_description_model

type ReminderDescription string

func NewReminderDescription(s string) ReminderDescription {
	return ReminderDescription(s)
}
