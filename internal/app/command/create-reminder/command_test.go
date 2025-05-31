package create_reminder_command

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	title, err := reminder_title_model.NewReminderTitle("title")
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription("description")
	require.NoError(t, err)

	command := NewCommand(title, description)
	require.Equal(t, title, command.GetReminderTitle())
	require.Equal(t, description, command.GetReminderDescription())
}

func TestCommand_Validate(t *testing.T) {
	t.Parallel()

	command := Command{
		title:       "title",
		description: "description",
	}
	err := command.Validate()
	require.NoError(t, err)
}

func TestNewCommand_Validate_Error(t *testing.T) {
	t.Parallel()

	t.Run("invalid title", func(t *testing.T) {
		command := Command{
			title:       "",
			description: "description",
		}
		err := command.title.Validate()
		require.Error(t, err)
	})

	t.Run("invalid description", func(t *testing.T) {
		t.Parallel()

		command := Command{
			title:       "title",
			description: "",
		}
		err := command.description.Validate()
		require.Error(t, err)
	})
}
