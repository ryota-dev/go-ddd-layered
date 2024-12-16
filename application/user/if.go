package user

import (
	"context"
	"go-ddd-layered/application/user/command"
)

type UserIF interface {
	Register(ctx context.Context, cmd command.RegisterUser) error
}
