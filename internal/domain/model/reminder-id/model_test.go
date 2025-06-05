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

	v := time.Now().Unix()

	reminderID, err := NewReminderID(v)
	require.NoError(t, err)
	require.Equal(t, v, int64(reminderID))
}

func TestReminderID_Validate(t *testing.T) {
	t.Parallel()

	v := time.Now().Unix()

	reminderID, err := NewReminderID(v)
	require.NoError(t, err)
	require.NoError(t, reminderID.Validate())
}

func TestReminderID_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderID := ReminderID(0)
	require.Error(t, reminderID.Validate())
}
