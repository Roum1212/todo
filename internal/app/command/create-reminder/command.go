package create_reminder_command

import (
	"fmt"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Command struct {
	title       reminder_title_model.ReminderTitle
	description reminder_description_model.ReminderDescription
}

func (x Command) GetReminderDescription() reminder_description_model.ReminderDescription {
	return x.description
}

func (x Command) GetReminderTitle() reminder_title_model.ReminderTitle {
	return x.title
}

func (x Command) Validate() error {
	if err := x.title.Validate(); err != nil {
		return fmt.Errorf("invalid title: %w", err)
	}

	if err := x.description.Validate(); err != nil {
		return fmt.Errorf("invalid description: %w", err)
	}

	return nil
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
