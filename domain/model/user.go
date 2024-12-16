package model

import (
	"go-ddd-layered/application/user/command"
	"go-ddd-layered/pkg/utils"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(cmd command.RegisterUser) (user User, err error) {
	utils.MarshalAndInsert(cmd, &user)
	return
}
