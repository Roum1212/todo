package get_all_reminders_http_handler

import (
	"encoding/json"
	"net/http"

	get_all_reminders_quer "github.com/Roum1212/todo/internal/app/query/get-all-reminders"
)

type Handler struct {
	queryHandler get_all_reminders_quer.Handler
}

const Endpoint = "/reminders/"

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reminder, err := x.queryHandler.Handle(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(reminder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

}

func NewHandler(commandHandler get_all_reminders_quer.Handler) Handler {
	return Handler{
		queryHandler: commandHandler,
	}
}
