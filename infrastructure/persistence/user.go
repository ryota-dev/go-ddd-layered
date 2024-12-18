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

func (up *userPersistence) List(ctx context.Context, page, perPage *int32) (users *[]model.User, err error) {
	var defaultPage, defaultPerPage int32 = 1, 10
	if page == nil {
		page = &defaultPage
	}
	if perPage == nil {
		perPage = &defaultPerPage
	}

	var dbUsers []dbmodel.User

	if err := utils.Paginator(up.db, *page, *perPage).Find(&dbUsers).Error; err != nil {
		return nil, err
	}

	utils.MarshalAndInsert(dbUsers, &users)

	return users, nil
}
