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
)

func TestHandler_ServeHTTP_Created(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := Request{
		Title:       "title",
		Description: "description",
	}

	commandHandlerMock := create_reminder_command.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c create_reminder_command.Command) {
			require.Equal(mc, request.Title, string(c.GetTitle()))
			require.Equal(t, request.Description, string(c.GetDescription()))
		}).
		Return(nil)

	httpHandler := NewHandler(commandHandlerMock)

	requestBody, _ := json.Marshal(request)

	r := httptest.NewRequest(http.MethodPost, Endpoint, bytes.NewBuffer(requestBody))

	r.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusCreated, recorder.Code)
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	commandHandlerMock := create_reminder_command.NewCommandHandlerMock(mc)

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
		requestBody, _ := json.Marshal(tt)

		r := httptest.NewRequest(http.MethodPost, Endpoint, bytes.NewBuffer(requestBody))

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

	commandHandlerMock := create_reminder_command.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Inspect(func(ctx context.Context, c create_reminder_command.Command) {
			require.Equal(mc, request.Title, string(c.GetTitle()))
			require.Equal(t, request.Description, string(c.GetDescription()))
		}).
		Return(assert.AnError)

	httpHandler := NewHandler(commandHandlerMock)

	requestBody, _ := json.Marshal(request)

	r := httptest.NewRequest(http.MethodPost, Endpoint, bytes.NewBuffer(requestBody))

	r.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	httpHandler.ServeHTTP(recorder, r)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
}
