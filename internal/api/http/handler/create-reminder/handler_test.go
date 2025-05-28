package create_reminder_http_handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	reminder_aggregate "github.com/Roum1212/todo/internal/domain/aggregate/reminder"
)

func TestHandler_ServeHTTP_Created(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	defer mc.Wait(time.Second)

	reminderRepositoryMock := reminder_aggregate.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, r reminder_aggregate.Reminder) {
			require.Equal(mc, "title", string(r.GetTitle()))
			require.Equal(mc, "description", string(r.GetDescription()))
		}).
		Return(nil)

	commandHandler := create_reminder_command.NewHandler(reminderRepositoryMock)

	httpHandler := NewHandler(commandHandler)

	body := []byte(`{"title":"title","description":"description"}`)
	req := httptest.NewRequest(http.MethodPost, Endpoint, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusCreated, recorder.Code)
}

func TestHandler_ServeHTTP_Bad_Request(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	defer mc.Wait(time.Second)

	reminderRepositoryMock := reminder_aggregate.NewReminderRepositoryMock(mc)

	commandHandler := create_reminder_command.NewHandler(reminderRepositoryMock)

	httpHandler := NewHandler(commandHandler)

	TableTest := []struct {
		body []byte
	}{
		{[]byte(`{"title":123,"description":"description"}`)},
		{[]byte(`{"title":"title only"`)},
		{[]byte(``)},
	}

	for _, tt := range TableTest {
		req := httptest.NewRequest(http.MethodPost, Endpoint, bytes.NewBuffer(tt.body))

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		httpHandler.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestHandler_ServeHTTP_Internal_Server_Error(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	defer mc.Wait(time.Second)

	err := errors.New("internal server error")

	reminderRepositoryMock := reminder_aggregate.NewReminderRepositoryMock(mc).
		SaveReminderMock.
		Inspect(func(ctx context.Context, r reminder_aggregate.Reminder) {
			require.Equal(mc, "title", string(r.GetTitle()))
			require.Equal(mc, "description", string(r.GetDescription()))
		}).
		Return(err)

	commandHandler := create_reminder_command.NewHandler(reminderRepositoryMock)

	httpHandler := NewHandler(commandHandler)

	body := []byte(`{"title":"title","description":"description"}`)
	req := httptest.NewRequest(http.MethodPost, Endpoint, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
