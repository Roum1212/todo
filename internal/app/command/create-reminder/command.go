package create_reminder_command

import (
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type CreateCommand struct {
	title       reminder_title_model.ReminderTitle
	description reminder_description_model.ReminderDescription
}

func NewCreateCommand(
	title reminder_title_model.ReminderTitle,
	description reminder_description_model.ReminderDescription,
) CreateCommand {
	return CreateCommand{
		title:       title,
		description: description,
	}
}
