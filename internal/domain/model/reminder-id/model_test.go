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
	reminderID, err := NewReminderID("1234567890") //nolint:wsl // OK.

	require.NoError(t, err)
	require.Equal(t, ReminderID(1234567890), reminderID)
}

func TestNewReminderID_Err(t *testing.T) {
	t.Parallel()
	reminderID, err := NewReminderID("?") //nolint:wsl // OK.

	require.Error(t, err)
	require.Zero(t, reminderID)
}
