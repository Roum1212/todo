package reminder_title_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewReminderTitle(t *testing.T) {
	t.Parallel()

	reminderTitle := NewReminderTitle("abc123")
	require.Equal(t, "abc123", string(reminderTitle))
}
