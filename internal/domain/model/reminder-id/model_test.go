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

	reminderID, err := NewReminderID("1234567890")
	require.NoError(t, err)
	require.Equal(t, 1234567890, int(reminderID))
}

func TestNewReminderID_Err(t *testing.T) {
	t.Parallel()

	reminderID, err := NewReminderID("?")
	require.Error(t, err)
	require.Zero(t, reminderID)
}
