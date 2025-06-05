package get_all_reminders_http_handler

import (
	"encoding/json"
	"errors"
	"net/http"

	get_all_reminders_quer "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
)

const Endpoint = "/reminders"

type Handler struct {
	queryHandler get_all_reminders_quer.QueryHandler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reminders, err := x.queryHandler.HandleQuery(r.Context())
	if err != nil {
		switch {
		case errors.Is(err, get_all_reminders_quer.ErrRemindersNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}

	reminderDTOs := NewReminderDTOs(reminders)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(reminderDTOs) //nolint:errcheck,errchkjson // OK.
}

func NewHandler(queryHandler get_all_reminders_quer.QueryHandler) Handler {
	return Handler{
		queryHandler: queryHandler,
	}
}
