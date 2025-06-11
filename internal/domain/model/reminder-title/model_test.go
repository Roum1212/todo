package reminder_title_model

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderTitle(t *testing.T) {
	t.Parallel()

	s := rand.Text()

	reminderTitle, err := NewReminderTitle(s)
	require.NoError(t, err)
	require.Equal(t, s, string(reminderTitle))
}

func TestNewReminderTitle_Error(t *testing.T) {
	t.Parallel()

	s := ""

	reminderTitle, err := NewReminderTitle(s)
	require.Error(t, err)
	require.Empty(t, reminderTitle)
}

func TestReminderTitle_Validate(t *testing.T) {
	t.Parallel()

	s := rand.Text()

	reminderTitle, err := NewReminderTitle(s)
	require.NoError(t, err)
	require.NoError(t, reminderTitle.Validate())
}

func TestReminderTitle_Validate_Error(t *testing.T) {
	t.Parallel()

	reminderTitle := ReminderTitle("")
	require.Error(t, reminderTitle.Validate())
}
