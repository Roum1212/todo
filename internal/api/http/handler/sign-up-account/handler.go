package sign_up_account_http_handler

import (
	"encoding/json"
	"net/http"

	sign_up_account_command "github.com/Roum1212/todo/internal/app/command/sign-up-account"
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

const Endpoint = "/sign-up-account"

type Handler struct {
	commandHandler sign_up_account_command.Handler
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
		sign_up_account_command.NewCommand(login, password),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHandler(commandHandler sign_up_account_command.Handler) Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}

type Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
