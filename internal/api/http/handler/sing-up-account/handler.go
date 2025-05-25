package sing_up_account_http_handler

import (
	"encoding/json"
	"net/http"

	sing_up_account_command "github.com/Roum1212/todo/internal/app/command/sing-up-account"
	account_login_model "github.com/Roum1212/todo/internal/domain/model/account-login"
	account_password_model "github.com/Roum1212/todo/internal/domain/model/account-password"
)

const Endpoint = "/sing-up-account"

type Handler struct {
	commandHandler sing_up_account_command.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request RequestToSingUp

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	accountLogin := account_login_model.NewAccountLogin(request.Login)
	accountPassword := account_password_model.NewAccountPassword(request.Password)

	if err := x.commandHandler.Handle(
		r.Context(),
		sing_up_account_command.NewCommand(accountLogin, accountPassword),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHandler(commandHandler sing_up_account_command.Handler) Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}

type RequestToSingUp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
