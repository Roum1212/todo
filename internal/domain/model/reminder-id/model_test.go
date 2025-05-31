package reminder_id_model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateReminderID(t *testing.T) {
	t.Parallel()

	reminderID1 := GenerateReminderID()

	time.Sleep(2 * time.Nanosecond)

	reminderID2 := GenerateReminderID()

	require.NotEqual(t, reminderID1, reminderID2)
}

func TestNewReminderID(t *testing.T) {
	t.Parallel()

	reminderID, err := NewReminderID("123")
	require.NoError(t, err)
	require.Equal(t, 123, int(reminderID))
}

func TestReminderID_Validate(t *testing.T) {
	t.Parallel()

	reminderID := ReminderID(123)
	err := reminderID.Validate()
	require.NoError(t, err)
}

func TestNewReminderID_Err(t *testing.T) {
	t.Parallel()

	_, err := NewReminderID("abc123")
	require.Error(t, err)
}

func TestReminderID_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderID := ReminderID(0)
	err := reminderID.Validate()
	require.Error(t, err)
}
