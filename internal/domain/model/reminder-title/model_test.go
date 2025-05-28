package reminder_title_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderTitle(t *testing.T) {
	t.Parallel()

	reminderTitle, err := NewReminderTitle("abc123")

	require.NoError(t, err)
	require.Equal(t, "abc123", string(reminderTitle))
}
