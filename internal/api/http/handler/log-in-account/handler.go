package log_in_account_http_handler

import (
	"encoding/json"
	"net/http"

	log_in_account_command "github.com/Roum1212/todo/internal/app/command/log-in-account"
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

const Endpoint = "/log-in-account"

type Handler struct {
	commandHandler log_in_account_command.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	login := account_login_model.NewAccountLogin(request.Login)
	password := account_password_model.NewAccountPassword(request.Password)

	if err := x.commandHandler.Handle(
		r.Context(),
		log_in_account_command.NewCommand(login, password),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewHandler(commandHandler log_in_account_command.Handler) Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}

type Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
