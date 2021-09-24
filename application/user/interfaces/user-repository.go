package interfaces

import (
	"weightexplorer/application/user/ports/dto"
	"weightexplorer/domain"
)

type UserRepository interface {
	Validate(user *domain.User) error
	Save(user *domain.User) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	GetById(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Update(id string, User *domain.User) (*domain.User, error)
	Delete(id string) error
	Patch(id string, user *dto.UserPatch) (*domain.User, error)
}
