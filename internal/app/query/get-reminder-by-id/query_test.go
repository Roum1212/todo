package get_reminder_by_id_query

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestNewQuery(t *testing.T) {
	t.Parallel()

	TableTest := []struct {
		reminderID string
	}{
		{reminderID: ""},
		{reminderID: "12345"},
		{reminderID: "-2 5 0 3 5"},
	}

	for _, tt := range TableTest {
		reminderID, _ := reminder_id_model.NewReminderID(tt.reminderID) //nolint:errcheck // OK.

		query := NewQuery(reminderID)
		require.Equal(t, reminderID, query.reminderID)
	}
}
