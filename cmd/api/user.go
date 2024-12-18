package api

import (
	"go-ddd-layered/application/user"
	"go-ddd-layered/infrastructure/persistence"
	"go-ddd-layered/interface/handler"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserEndpoints(r *mux.Router, db *gorm.DB) {
	userPersistence := persistence.NewUserPersistence(db)
	userUsecase := user.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	r.HandleFunc("/users", userHandler.HandleRegister).Methods("POST")
	r.HandleFunc("/users", userHandler.HandleList).Methods("GET")
}
