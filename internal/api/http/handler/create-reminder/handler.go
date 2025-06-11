package create_reminder_http_handler

import (
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

const Endpoint = "/reminders"

const tracerName = "github.com/Roum1212/todo/internal/api/http/handler/create-reminder/handler.go"

type Handler struct {
	commandHandler create_reminder_command.CommandHandler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reminderTitle, err := reminder_title_model.NewReminderTitle(request.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reminderDescription, err := reminder_description_model.NewReminderDescription(request.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = x.commandHandler.HandleCommand(
		r.Context(),
		create_reminder_command.NewCommand(reminderTitle, reminderDescription),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHTTPHandler(commandHandler create_reminder_command.CommandHandler) http.Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type tracerHandler struct {
	handler http.Handler
	tracer  trace.Tracer
}

func (x tracerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, span := x.tracer.Start(r.Context(), "Handler.ServeHTTP")
	defer span.End()

	x.handler.ServeHTTP(w, r)
}

func NewHTTPHandlerWithTracer(handler http.Handler) http.Handler {
	tracer := otel.Tracer(tracerName)

	return tracerHandler{
		handler: handler,
		tracer:  tracer,
	}
}
