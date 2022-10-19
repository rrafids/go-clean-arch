package postgresql

import (
	"go-clean-arch/domain/model"
	presenter "go-clean-arch/usecase/presenter/postgresql"
	repository "go-clean-arch/usecase/repository/postgresql"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	GetAll(users []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) GetAll(users []*model.User) ([]*model.User, error) {
	users, err := us.UserRepository.FindAll(users)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(users), nil
}
