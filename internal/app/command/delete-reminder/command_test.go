package delete_reminder_command

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	reminderID, err := reminder_id_model.NewReminderID("123")

	require.NoError(t, err)

	command := NewCommand(reminderID)

	require.Equal(t, reminderID, command.reminderID)
}
