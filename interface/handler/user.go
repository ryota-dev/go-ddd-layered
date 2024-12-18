package handler

import (
	"go-ddd-layered/application/user"
	"go-ddd-layered/application/user/command"
	"go-ddd-layered/interface/response"
	"net/http"
	"strconv"
)

type UserHandler interface {
	HandleRegister(http.ResponseWriter, *http.Request)
	HandleList(http.ResponseWriter, *http.Request)
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

func (uh *userHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("perPage")

	var page, perPage *int32

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			pageValue := int32(p)
			page = &pageValue
		}
	}

	if perPageStr != "" {
		if p, err := strconv.Atoi(perPageStr); err == nil {
			perPageValue := int32(p)
			perPage = &perPageValue
		}
	}

	cmd := command.ListUser{
		Page:    page,
		PerPage: perPage,
	}

	users, err := uh.userUsecase.List(r.Context(), cmd)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, "failed to list users")
		return
	}

	response.Respond(w, http.StatusOK, users)
}
