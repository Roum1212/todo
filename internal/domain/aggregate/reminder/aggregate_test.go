package reminder_aggregate

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestReminder(t *testing.T) {
	t.Parallel()

	id, err := reminder_id_model.NewReminderID(time.Now().Unix())
	require.NoError(t, err)

	title, err := reminder_title_model.NewReminderTitle(rand.Text())
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription("abc")
	require.NoError(t, err)

	reminder := NewReminder(id, title, description)
	require.Equal(t, id, reminder.GetID())
	require.Equal(t, title, reminder.GetTitle())
	require.Equal(t, description, reminder.GetDescription())
}
