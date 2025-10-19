package usecases

import (
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/entities"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/repository"
)

type GetUserById struct {
	repo repository.UserRepository
}

func NewGetUserById(repo *dao.MySQLUserDao) *GetUserById {
	return &GetUserById{
		repo: repo,
	}
}

func (u *GetUserById) Execute(id int) (entities.User, error) {
	user, err := u.repo.GetUserById(id)
	return user, err
}
