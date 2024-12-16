package user

import (
	"context"
	"go-ddd-layered/application/user/command"
	"go-ddd-layered/domain/model"
	"go-ddd-layered/domain/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserIF {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) Register(ctx context.Context, cmd command.RegisterUser) error {
	user, err := model.NewUser(cmd)
	if err != nil {
		return err
	}

	if err := uu.userRepository.Register(ctx, user); err != nil {
		return err
	}

	return nil
}