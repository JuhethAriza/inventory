package controllers
<<<<<<< HEAD

import (
	"strconv"

	r "github.com/JuhethAriza/inventory/src/common/response"
	"github.com/JuhethAriza/inventory/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetProductByIDController struct {
	usecase *usecases.GetProductByID
	result  *r.Result
}

func NewGetProductByIDController(uc *usecases.GetProductByID, r *r.Result) *GetProductByIDController {
	return &GetProductByIDController{usecase: uc, result: r}
}

func (c *GetProductByIDController) Run(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.result.Bad(ctx, "ID inválido")
	}

	product, err := c.usecase.Execute(id)
	if err != nil {
		return c.result.Bad(ctx, err.Error())
	}

	return c.result.Ok(ctx, product)
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
