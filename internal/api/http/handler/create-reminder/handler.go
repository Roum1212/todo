package create_reminder_http_handler

import (
	"encoding/json"
	"net/http"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

const Endpoint = "/reminders"

type Handler struct {
	commandHandler create_reminder_command.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request RequestToSave

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	title, err := reminder_title_model.NewReminderTitle(request.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	description, err := reminder_description_model.NewReminderDescription(request.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := x.commandHandler.Handle(
		r.Context(),
		create_reminder_command.NewCommand(title, description),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHandler(commandHandler create_reminder_command.Handler) Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}

type RequestToSave struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
