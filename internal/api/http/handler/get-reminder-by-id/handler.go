package get_reminder_by_id_http_handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	get_reminder_by_id_quary "github.com/Roum1212/todo/internal/app/query/get-reminder-by-id"
	reminder_id_model "github.com/Roum1212/todo/internal/domain/model/reminder-id"
)

const Endpoint = "/reminders/" + paramID

const paramID = ":id"

type Handler struct {
	queryHandler get_reminder_by_id_quary.QueryHandler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	reminderID, err := reminder_id_model.NewReminderIDFromString(
		params.ByName(strings.TrimPrefix(paramID, ":")),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reminder, err := x.queryHandler.HandleQuery(
		r.Context(),
		get_reminder_by_id_quary.NewQuery(reminderID),
	)
	if err != nil {
		switch {
		case errors.Is(err, get_reminder_by_id_quary.ErrReminderNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}

	reminderDTO := NewReminder(
		reminder.GetID(),
		reminder.GetTitle(),
		reminder.GetDescription(),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(reminderDTO) //nolint:errcheck,errchkjson // OK.
}

func NewHandler(queryHandler get_reminder_by_id_quary.QueryHandler) Handler {
	return Handler{
		queryHandler: queryHandler,
	}
}
