package controllers
<<<<<<< HEAD

import (
	r "github.com/JuhethAriza/inventory/src/common/response"
	"github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
	usecases "github.com/JuhethAriza/inventory/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type CreateProductController struct {
	usecase *usecases.CreateProduct
	result  *r.Result
}

func NewCreateProductController(u *usecases.CreateProduct, res *r.Result) *CreateProductController {
	return &CreateProductController{
		usecase: u,
		result:  res,
	}
}

func (pc *CreateProductController) Run(c *fiber.Ctx) error {
	var req dto.ProductDTO
	if err := c.BodyParser(&req); err != nil {
		return pc.result.Error(c, err.Error())
	}
	product, err := pc.usecase.Execute(req)
	if err != nil {
		return pc.result.Error(c, err.Error())
	}

	return pc.result.Ok(c, product)
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
