package persistence

import (
	"context"
	"go-ddd-layered/domain/model"
	"go-ddd-layered/domain/repository"
	"go-ddd-layered/infrastructure/dbmodel"
	"go-ddd-layered/pkg/utils"

	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{
		db: db,
	}
}

func (up *userPersistence) Register(ctx context.Context, user model.User) error {
	var dbUser dbmodel.User
	utils.MarshalAndInsert(user, &dbUser)

	if result := up.db.Debug().Create(&dbUser); result.Error != nil {
		return result.Error
	}

	return nil
}
