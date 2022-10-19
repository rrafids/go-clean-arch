package postgresql

import "go-clean-arch/domain/model"

type UserPresenter interface {
	ResponseUsers(users []*model.User) []*model.User
}
