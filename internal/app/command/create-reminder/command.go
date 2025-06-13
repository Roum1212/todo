package create_reminder_command

import (
	"fmt"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

type Command struct {
	id          reminder_id_model.ReminderID
	title       reminder_title_model.ReminderTitle
	description reminder_description_model.ReminderDescription
}

func (x Command) GetID() reminder_id_model.ReminderID {
	return x.id
}

func (x Command) GetTitle() reminder_title_model.ReminderTitle {
	return x.title
}

func (x Command) GetDescription() reminder_description_model.ReminderDescription {
	return x.description
}

func (x Command) Validate() error {
	if err := x.id.Validate(); err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	if err := x.title.Validate(); err != nil {
		return fmt.Errorf("invalid title: %w", err)
	}

	if err := x.description.Validate(); err != nil {
		return fmt.Errorf("invalid description: %w", err)
	}

	return nil
}

func NewCommand(
	id reminder_id_model.ReminderID,
	title reminder_title_model.ReminderTitle,
	description reminder_description_model.ReminderDescription,
) Command {
	return Command{
		id:          id,
		title:       title,
		description: description,
	}
}
