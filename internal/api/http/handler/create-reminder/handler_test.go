package create_reminder_http_handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	"github.com/Roum1212/todo/internal/app/command/create-reminder/mock"
)

func TestHandler_ServeHTTP_Created(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := Request{
		Title:       "title",
		Description: "description",
	}

	commandHandlerMock := mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c create_reminder_command.Command) {
			require.Equal(mc, request.Title, string(c.GetReminderTitle()))
			require.Equal(t, request.Description, string(c.GetReminderDescription()))
		}).
		Return(nil)

	httpHandler := NewHandler(commandHandlerMock)

	requestBody, err := json.Marshal(request) //nolint:errchkjson // OK.
	require.NoError(t, err)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		Endpoint,
		bytes.NewReader(requestBody),
	)
	r.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusCreated, recorder.Code)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	commandHandlerMock := mock.NewCommandHandlerMock(mc)

	httpHandler := NewHandler(commandHandlerMock)

	tests := []Request{
		{
			Title:       "",
			Description: "",
		},
		{
			Title:       "title",
			Description: "",
		},
		{
			Title:       "",
			Description: "description",
		},
	}

	for _, tt := range tests {
		requestBody, err := json.Marshal(tt) //nolint:errchkjson // OK.
		require.NoError(t, err)

		r := httptest.NewRequestWithContext(
			t.Context(),
			http.MethodPost,
			Endpoint,
			bytes.NewReader(requestBody),
		)
		r.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		httpHandler.ServeHTTP(recorder, r)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := Request{
		Title:       "title",
		Description: "description",
	}

	commandHandlerMock := mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c create_reminder_command.Command) {
			require.Equal(mc, request.Title, string(c.GetReminderTitle()))
			require.Equal(t, request.Description, string(c.GetReminderDescription()))
		}).
		Return(assert.AnError)

	httpHandler := NewHandler(commandHandlerMock)

	requestBody, err := json.Marshal(request) //nolint:errchkjson // OK.
	require.NoError(t, err)

	r := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		Endpoint,
		bytes.NewReader(requestBody),
	)
	r.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
