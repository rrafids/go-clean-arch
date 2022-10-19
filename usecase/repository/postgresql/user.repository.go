package postgresql

import "go-clean-arch/domain/model"

type UserRepository interface {
	FindAll(users []*model.User) ([]*model.User, error)
}
