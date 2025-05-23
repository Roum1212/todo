package reminder_id_model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require" //nolint:depguard // OK.
)

func TestGenerateReminderID(t *testing.T) {
	t.Parallel()

	id1 := GenerateReminderID()

	time.Sleep(1 * time.Second)

	id2 := GenerateReminderID()

	require.NotEqual(t, id1, id2, "Reminder ID should not be the same")
}

func TestGenerateReminderIDValid(t *testing.T) {
	t.Parallel()

	text, err := NewReminderID("1234567890")

	require.NoError(t, err)
	require.Equal(t, ReminderID(1234567890), text)
}

func TestGenerateReminderIDInvalid(t *testing.T) {
	t.Parallel()

	str := "abc"
	_, err := NewReminderID(str)

	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid syntax")
}
