package controllers

import (
	"strconv"

	r "github.com/JuhethAriza/inventory/src/common/response"
	dto "github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
	usecases "github.com/JuhethAriza/inventory/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type UpdateProductController struct {
	usecase *usecases.UpdateProduct
	result  *r.Result
}

func NewUpdateProductController(uc *usecases.UpdateProduct, res *r.Result) *UpdateProductController {
	return &UpdateProductController{usecase: uc, result: res}
}

func (c *UpdateProductController) Run(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.result.Bad(ctx, "ID inválido")
	}

	var req dto.UpdateProductDTO
	if err := ctx.BodyParser(&req); err != nil {
		return c.result.Bad(ctx, "Error al parsear el cuerpo de la solicitud")
	}

	resp, err := c.usecase.Execute(id, req)
	if err != nil {
		return c.result.Bad(ctx, err.Error())
	}

	return c.result.Ok(ctx, resp)
}
