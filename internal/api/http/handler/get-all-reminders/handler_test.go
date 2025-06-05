package get_all_reminders_http_handler

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	get_all_reminders_query "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	"github.com/Roum1212/todo/internal/app/query/get-all-reminders/mock"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestHandler_ServeHTTP_OK(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminder1 := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)
	reminder2 := reminder_aggregate.NewReminder(
		reminder_id_model.GenerateReminderID(),
		reminder_title_model.ReminderTitle(rand.Text()),
		reminder_description_model.ReminderDescription(rand.Text()),
	)
	reminders := []reminder_aggregate.Reminder{reminder1, reminder2}

	queryHandlerMock := get_all_reminders_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext).
		Return(reminders, nil)

	httpHandler := NewHTTPHandler(queryHandlerMock)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		Endpoint,
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	reminderDTOs := NewReminderDTOs(reminders)

	var gotReminderDTOs []Reminder

	require.NoError(t, json.NewDecoder(recorder.Body).Decode(&gotReminderDTOs))
	require.Equal(t, gotReminderDTOs, reminderDTOs)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := get_all_reminders_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext).
		Return(nil, assert.AnError)

	httpHandler := NewHTTPHandler(queryHandlerMock)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		Endpoint,
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func TestHandler_ServeHTTP_StatusNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := get_all_reminders_query_mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Expect(minimock.AnyContext).
		Return(nil, get_all_reminders_query.ErrReminderNotFound)

	httpHandler := NewHTTPHandler(queryHandlerMock)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		Endpoint,
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusNotFound, recorder.Code)
}
