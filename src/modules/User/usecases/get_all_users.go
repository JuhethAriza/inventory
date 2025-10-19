package usecases
<<<<<<< HEAD

import (
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/entities"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/repository"
)

type GetAllUsers struct {
	repo repository.UserRepository
}

func NewGetAllUsers(repo *dao.MySQLUserDao) *GetAllUsers {
	return &GetAllUsers{
		repo: repo,
	}
}

func (u *GetAllUsers) Execute() ([]entities.User, error) {
	users, err := u.repo.GetAllUsers(1, 10)
	return users, err
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
