package create_reminder_http_handler

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	"github.com/Roum1212/todo/internal/app/command/create-reminder/mock"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

func TestHandler_ServeHTTP_Created(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := Request{
		Title:       rand.Text(),
		Description: rand.Text(),
	}

	title, err := reminder_title_model.NewReminderTitle(request.Title)
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(request.Description)
	require.NoError(t, err)

	commandHandlerMock := create_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, create_reminder_command.NewCommand(title, description)).
		Return(nil)

	httpHandler := NewHTTPHandler(commandHandlerMock)

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

	t.Run("empty title", func(t *testing.T) {
		t.Parallel()

		httpHandler := NewHTTPHandler(nil)

		request := Request{
			Title:       "",
			Description: rand.Text(),
		}

		requestBody, err := json.Marshal(request)
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
	})

	t.Run("empty description", func(t *testing.T) {
		t.Parallel()

		httpHandler := NewHTTPHandler(nil)

		request := Request{
			Title:       rand.Text(),
			Description: "",
		}

		requestBody, err := json.Marshal(request)
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
	})

	t.Run("invalid request body", func(t *testing.T) {
		t.Parallel()

		requestBody := []byte(`{"title": "title", "description": "description"`)

		httpHandler := NewHTTPHandler(nil)

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
	})
}

func TestHandler_ServeHTTP_InternalServerError(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	request := Request{
		Title:       rand.Text(),
		Description: rand.Text(),
	}

	title, err := reminder_title_model.NewReminderTitle(request.Title)
	require.NoError(t, err)

	description, err := reminder_description_model.NewReminderDescription(request.Description)
	require.NoError(t, err)

	commandHandlerMock := create_reminder_command_mock.NewCommandHandlerMock(mc).
		HandleCommandMock.
		Expect(minimock.AnyContext, create_reminder_command.NewCommand(title, description)).
		Return(assert.AnError)

	httpHandler := NewHTTPHandler(commandHandlerMock)

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
