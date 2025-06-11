package get_reminder_by_id_query

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestNewQuery(t *testing.T) {
	t.Parallel()

	reminderID := reminder_id_model.GenerateReminderID()

	query := NewQuery(reminderID)
	require.Equal(t, reminderID, query.GetID())
}

func TestQuery_Validate(t *testing.T) {
	t.Parallel()

	query := NewQuery(reminder_id_model.GenerateReminderID())
	require.NoError(t, query.Validate())
}

func TestQuery_Validate_Error(t *testing.T) {
	t.Parallel()

	query := NewQuery(0)
	require.Error(t, query.Validate())
}
