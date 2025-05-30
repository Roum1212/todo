package get_reminder_by_id_query

import (
	"testing"

	"github.com/stretchr/testify/require"

	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

func TestNewQuery(t *testing.T) {
	t.Parallel()

	reminderID, err := reminder_id_model.NewReminderID("123")
	require.NoError(t, err)

	query := NewQuery(reminderID)
	require.Equal(t, reminderID, query.GetReminderID())
}

func TestQuery_Validate(t *testing.T) {
	t.Parallel()

	query := Query{
		reminderID: 123,
	}
	err := query.Validate()
	require.NoError(t, err)
}

func TestQuery_Validate_Error(t *testing.T) {
	t.Parallel()

	query := Query{
		reminderID: 0,
	}
	err := query.Validate()
	require.Error(t, err)
}
