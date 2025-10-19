package usecases
<<<<<<< HEAD

import (
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/entities"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type CreateUsers struct {
	repo repository.UserRepository
}

func NewCreateUsers(repo *dao.MySQLUserDao) *CreateUsers {
	return &CreateUsers{
		repo: repo,
	}
}

func (u *CreateUsers) Execute(user *entities.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return u.repo.CreateUser(user)
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
