package create_reminder_command

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	id := reminder_id_model.GenerateReminderID()

	command := NewCommand(id, title, description)
	require.Equal(t, id, command.GetID())
	require.Equal(t, title, command.GetTitle())
	require.Equal(t, description, command.GetDescription())
}

func TestCommand_Validate(t *testing.T) {
	t.Parallel()

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(rand.Text())
	require.NoError(t, err)

	command := NewCommand(reminder_id_model.GenerateReminderID(), title, description)
	require.NoError(t, command.Validate())
}

func TestCommand_Validate_Error(t *testing.T) {
	t.Parallel()

	t.Run("invalid id", func(t *testing.T) {
		t.Parallel()

		title, err := reminder_title_model.NewReminderTitle(rand.Text())
		require.NoError(t, err)

		description, err := reminder_description_model.NewReminderDescription(rand.Text())
		require.NoError(t, err)

		command := NewCommand(0, title, description)
		require.Error(t, command.Validate())
	})

	t.Run("invalid title", func(t *testing.T) {
		t.Parallel()

		description, err := reminder_description_model.NewReminderDescription(rand.Text())
		require.NoError(t, err)

		command := NewCommand(reminder_id_model.GenerateReminderID(), "", description)
		require.Error(t, command.Validate())
	})

	t.Run("invalid description", func(t *testing.T) {
		t.Parallel()

		title, err := reminder_title_model.NewReminderTitle(rand.Text())
		require.NoError(t, err)

		command := NewCommand(reminder_id_model.GenerateReminderID(), title, "")
		require.Error(t, command.Validate())
	})
}
