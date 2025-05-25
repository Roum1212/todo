package user_registration_handler_http

import (
	"encoding/json"
	"net/http"

	user_registration_command "github.com/Roum1212/todo/internal/app/command/user-registration"
	user_password_model "github.com/Roum1212/todo/internal/domain/model/user-password"
	user_login_model "github.com/Roum1212/todo/internal/domain/model/usre-login"
)

const Endpoint = "/user-registration"

type Handler struct {
	commandHandler user_registration_command.Handler
}

func (x Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request RequestToRegistration

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	userLogin := user_login_model.NewUserLogin(request.Login)
	userPassword := user_password_model.NewUserPassword(request.Password)

	if err := x.commandHandler.Handle(
		r.Context(),
		user_registration_command.NewCommand(userLogin, userPassword),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewHandler(commandHandler user_registration_command.Handler) Handler {
	return Handler{
		commandHandler: commandHandler,
	}
}

type RequestToRegistration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
