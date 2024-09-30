package repository

import (
	"errors"
	"gorm.io/gorm"
	"jamal/api/models/domain"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (repository AuthRepositoryImpl) Register(tx *gorm.DB, user domain.User) error {
	// <-- check acc already ?
	var existUser domain.User
	if err := tx.Where("user_name = ?", user.UserName).First(&existUser).Error; err == nil {
		return errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("failed to query database")
	}

	// <-- create
	result := tx.Create(&user).Error
	if result != nil {
		return errors.New("Error Create")
	}
	return nil
}

func (repository AuthRepositoryImpl) Login(tx *gorm.DB, user domain.User) error {
	var dataDB domain.User
	result := tx.Take(&dataDB, "user_name=? AND password=? ", user.UserName, user.Password).Error
	if result != nil {
		return errors.New("id Or password not wrong")
	}

	return nil
}

func (repository AuthRepositoryImpl) Logout(tx *gorm.DB, user domain.User) error {
	//TODO implement me
	panic("implement me")
}
