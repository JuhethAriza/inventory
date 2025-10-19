package repository
<<<<<<< HEAD

import "github.com/JuhethAriza/inventory/src/modules/User/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindByEmail(email string) (entities.User, error)
	GetAllUsers(page int, pageSize int) ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
