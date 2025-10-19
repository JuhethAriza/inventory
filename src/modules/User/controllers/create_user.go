package controllers
<<<<<<< HEAD

import (
	"github.com/JuhethAriza/inventory/src/modules/User/domain/entities"
	"github.com/JuhethAriza/inventory/src/modules/User/usecases"

	r "github.com/JuhethAriza/inventory/src/common/response"
	"github.com/gofiber/fiber/v2"
)

type CreateUsersController struct {
	usecase *usecases.CreateUsers
	result  *r.Result
}

func NewCreateUsersController(usecase *usecases.CreateUsers, r *r.Result) *CreateUsersController {
	return &CreateUsersController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *CreateUsersController) Run(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return ph.result.Bad(c, "Error al parsear el cuerpo de la solicitud")
	}

	if user.Email == "" {
		return ph.result.Bad(c, "El email es obligatorio")
	}

	if user.Password == "" {
		return ph.result.Bad(c, "La contraseña es obligatoria")
	}

	if len(user.Password) < 6 {
		return ph.result.Bad(c, "La contraseña debe tener al menos 6 caracteres")
	}

	if err := ph.usecase.Execute(&user); err != nil {
		return ph.result.Error(c, err)
	}

	return ph.result.Ok(c, "Usuario creado exitosamente")
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
