package user

import (
	"context"
	"errors"
	"go-ddd-layered/application/user/command"
	"go-ddd-layered/domain/model"
	"go-ddd-layered/domain/repository"
	"go-ddd-layered/pkg/utils"
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
	if cmd.Email == "" || cmd.Name == "" {
		return errors.New("invalid input")
	}

	user, err := model.NewUser(cmd)
	if err != nil {
		return err
	}

	if err := uu.userRepository.Register(ctx, user); err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) List(ctx context.Context, cmd command.ListUser) (usersDto *[]UserDTO, err error) {
	users, err := uu.userRepository.List(ctx, cmd.Page, cmd.PerPage)
	if err != nil {
		return nil, err
	}

	utils.MarshalAndInsert(users, &usersDto)

	return usersDto, nil
}
