package handler

import (
	"go-ddd-layered/application/user"
	"go-ddd-layered/application/user/command"
	"go-ddd-layered/interface/response"
	"net/http"
)

type UserHandler interface {
	HandleRegister(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUsecase user.UserIF
}

func NewUserHandler(uu user.UserIF) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

func (uh *userHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	cmd, err := command.NewRegisterUser(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, "failed to parse json")
		return
	}

	if err := uh.userUsecase.Register(r.Context(), cmd); err != nil {
		response.Error(w, http.StatusInternalServerError, err, "failed to register user")
		return
	}

	response.Respond(w, http.StatusOK, "user registered successfully")

}
