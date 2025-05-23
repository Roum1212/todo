package reminder_id_model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateReminderID(t *testing.T) {
	id1 := GenerateReminderID()
	time.Sleep(1 * time.Second)
	id2 := GenerateReminderID()

	require.NotEqual(t, id1, id2, "Reminder ID should not be the same")
}

func TestGenerateReminderIDValid(t *testing.T) {
	text, err := NewReminderID("1234567890")

	require.NoError(t, err)
	require.Equal(t, ReminderID(1234567890), text)
}

func TestGenerateReminderIDInvalid(t *testing.T) {
	str := "abc"
	_, err := NewReminderID(str)

	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid syntax")
}
