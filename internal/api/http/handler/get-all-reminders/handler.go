package get_all_reminders_http_handler

import (
	"encoding/json"
	"net/http"

	get_all_reminders_quer "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
	postgresql_reminder_repository "github.com/Roum1212/todo/internal/infra/repository/reminder/postgresql"
)

type Handler struct {
	queryHandler get_all_reminders_quer.Handler
}

const Endpoint = "/reminders"

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reminders, err := x.queryHandler.Handle(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reminderSlice := postgresql_reminder_repository.NewReminderJsonSlice(reminders)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(reminderSlice)
}

func NewHandler(queryHandler get_all_reminders_quer.Handler) Handler {
	return Handler{
		queryHandler: queryHandler,
	}
}
