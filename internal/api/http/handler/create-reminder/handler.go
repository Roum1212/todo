package create_reminder_http_handler

import (
	"encoding/json"
	"net/http"

	create_reminder_command "github.com/Roum1212/todo/internal/app/command/create-reminder"
	reminder_description_model "github.com/Roum1212/todo/internal/domain/model/reminder-description"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
	reminder_title_model "github.com/Roum1212/todo/internal/domain/model/reminder-title"
)

const Endpoint = "/reminders"

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

	reminderID := reminder_id_model.GenerateReminderID()

	if err = x.commandHandler.HandleCommand(
		r.Context(),
		create_reminder_command.NewCommand(reminderID, reminderTitle, reminderDescription),
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
