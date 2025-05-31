package get_reminder_by_id_http_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	"github.com/Roum1212/todo/internal/app/query/get-reminder-by-id/mock"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestHandler_ServeHTTP_OK(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	reminderID := reminder_id_model.ReminderID(123)
	reminderTitle := reminder_title_model.ReminderTitle("title")
	reminderDescription := reminder_description_model.ReminderDescription("description")
	reminder := reminder_aggregate.NewReminder(reminderID, reminderTitle, reminderDescription)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Inspect(func(ctx context.Context, q get_reminder_by_id_quary.Query) {
			require.Equal(t, reminderID, q.GetReminderID())
		}).
		Return(reminder, nil)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.Replace(Endpoint, ParamsID, "123", -1),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expectedBody := NewReminder(
		reminder.GetID(),
		reminder.GetTitle(),
		reminder.GetDescription())

	var gotBody Reminder
	err := json.NewDecoder(recorder.Body).Decode(&gotBody)
	require.NoError(t, err)
	require.Equal(t, expectedBody, gotBody)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := mock.NewQueryHandlerMock(mc)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.Replace(Endpoint, ParamsID, "abc", -1),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandler_ServeHTTP_StatusNotFound(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Inspect(func(ctx context.Context, q get_reminder_by_id_quary.Query) {
			require.Equal(t, reminder_id_model.ReminderID(123), q.GetReminderID())
		}).
		Return(reminder_aggregate.Reminder{}, get_reminder_by_id_quary.ErrReminderNotFound)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.Replace(Endpoint, ParamsID, "123", -1),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	queryHandlerMock := mock.NewQueryHandlerMock(mc).
		HandleQueryMock.
		Inspect(func(ctx context.Context, q get_reminder_by_id_quary.Query) {
			require.Equal(t, reminder_id_model.ReminderID(123), q.GetReminderID())
		}).
		Return(reminder_aggregate.Reminder{}, assert.AnError)

	httpHandler := NewHandler(queryHandlerMock)

	router := httprouter.New()
	router.Handler(http.MethodGet, Endpoint, httpHandler)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		strings.Replace(Endpoint, ParamsID, "123", -1),
		http.NoBody,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
