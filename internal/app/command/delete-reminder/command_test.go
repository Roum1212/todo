package delete_reminder_command

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	reminderID := reminder_id_model.GenerateReminderID()

	command := NewCommand(reminderID)
	require.Equal(t, reminderID, command.reminderID)
}

func TestCommand_Validate(t *testing.T) {
	t.Parallel()

	command := NewCommand(reminder_id_model.GenerateReminderID())
	require.NoError(t, command.Validate())
}

func TestCommand_Validate_Error(t *testing.T) {
	t.Parallel()

	command := NewCommand(0)
	require.Error(t, command.Validate())
}
