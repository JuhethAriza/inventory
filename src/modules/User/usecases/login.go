package usecases

import (
	"errors"

	auth "github.com/JuhethAriza/inventory/src/common/auth"
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/dto"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/entities"
	"github.com/JuhethAriza/inventory/src/modules/User/domain/repository"
	"github.com/JuhethAriza/inventory/src/modules/User/utils"
)

type LoginUser struct {
	repo repository.UserRepository
}

func NewLoginUser(repo *dao.MySQLUserDao) *LoginUser {
	return &LoginUser{repo: repo}
}

func (l *LoginUser) Execute(payload dto.LoginDTO) (string, *entities.User, error) {
	user, err := l.repo.FindByEmail(payload.Email)
	if err != nil {
		return "", nil, errors.New("usuario no encontrado")
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return "", nil, errors.New("contraseña inválida")
	}

	token, err := auth.GenerateToken(int(user.ID), user.Email)
	if err != nil {
		return "", nil, errors.New("error generando el token")
	}

	user.Password = ""

	return token, &user, nil
}
