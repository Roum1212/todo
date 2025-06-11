package reminder_id_model

import (
	"crypto/rand"
	"strconv"
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

func TestNewReminderID_Error(t *testing.T) {
	t.Parallel()

	reminderID, err := NewReminderID(0)
	require.Error(t, err)
	require.Empty(t, reminderID)
}

func TestReminderID_Validate(t *testing.T) {
	t.Parallel()

	reminderID, err := NewReminderID(time.Now().Unix())
	require.NoError(t, err)
	require.NoError(t, reminderID.Validate())
}

func TestReminderID_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderID := ReminderID(0)
	require.Error(t, reminderID.Validate())
}

func TestReminderIDFromString(t *testing.T) {
	t.Parallel()

	n := int64(123)

	reminderID, err := NewReminderIDFromString(strconv.FormatInt(n, 10))
	require.NoError(t, err)
	require.Equal(t, n, int64(reminderID))
}

func TestReminderIDFromString_Error(t *testing.T) {
	t.Parallel()

	reminderID, err := NewReminderIDFromString(rand.Text())
	require.Error(t, err)
	require.Empty(t, reminderID)
}
