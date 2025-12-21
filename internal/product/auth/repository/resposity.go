package repository

import (
	"WoodInspection/internal/product/auth/model"

	"gorm.io/gorm"
)

type UserReader interface {
	GetUserById(id int) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
}

type UserRepository interface {
	UserReader
}

type GormUserRepository struct {
	gormDB *gorm.DB
}

func NewGormUserRepository(gormDB *gorm.DB) UserRepository {
	return &GormUserRepository{gormDB: gormDB}
}

func (u GormUserRepository) GetUserById(id int) (*model.User, error) {
	user := &model.User{}
	err := u.gormDB.First(user, "id = ?", id).Error
	return user, err
}

func (u GormUserRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := u.gormDB.First(user, "username = ?", username).Error
	return user, err
}
