package interfaces

import (
	"weightexplorer/application/user/ports/viewmodel"
	"weightexplorer/domain"
)

type UserPresenter interface {
	UserResponse(userVM *viewmodel.UserVM) *domain.User
	UsersResponse(usersVM []*viewmodel.UserVM) []*domain.User
	UserVMResponse(userDom *domain.User) *viewmodel.UserVM
	UsersVMResponse(usersDom []*domain.User) []*viewmodel.UserVM
}
