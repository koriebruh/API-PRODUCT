package repository

import (
	"jamal/api/models/domain"
)

type UserRepository interface {
	Create(user domain.User) error
	FindByID(userID int) (domain.User, error)
	FindByUsername(username string) (domain.User, error)
	Update(user domain.User) error
	Delete(userID int) error
}
