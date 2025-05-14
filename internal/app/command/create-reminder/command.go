package create_reminder_command

import (
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Command struct {
	title       reminder_title_model.ReminderTitle
	description reminder_description_model.ReminderDescription
}

func NewCommand(
	title reminder_title_model.ReminderTitle,
	description reminder_description_model.ReminderDescription,
) Command {
	return Command{
		title:       title,
		description: description,
	}
}
