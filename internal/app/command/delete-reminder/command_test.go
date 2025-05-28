package delete_reminder_command

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()

	TableTest := []struct {
		reminderID string
	}{
		{reminderID: "1234"},
		{reminderID: ""},
		{reminderID: "-2 1 4 3"},
	}

	for _, tt := range TableTest {
		reminderID, _ := reminder_id_model.NewReminderID(tt.reminderID) //nolint:errcheck // OK.

		command := NewCommand(reminderID)
		require.Equal(t, reminderID, command.reminderID)
	}
}
