package repository

import (
	"context"
	"go-ddd-layered/domain/model"
)

type UserRepository interface {
	Register(ctx context.Context, user model.User) error
}
