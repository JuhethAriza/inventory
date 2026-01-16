package usecases

import (
	"fmt"

	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	dto "github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
	repository "github.com/JuhethAriza/inventory/src/modules/Producto/domain/repository"
	"github.com/JuhethAriza/inventory/src/modules/Producto/utils"
	"gorm.io/gorm"
)

type UpdateProduct struct {
	repo repository.ProductRepository
}

func NewUpdateProduct(repo *dao.MySQLProductDao) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (uc *UpdateProduct) Execute(id int, req dto.UpdateProductDTO) (*dto.ProductResponse, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID inválido")
	}

	if err := utils.ValidateUpdateProduct(req); err != nil {
		return nil, err
	}

	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("producto no encontrado")
		}
		return nil, fmt.Errorf("error al buscar el producto: %v", err)
	}

	product.CodigoProducto = req.CodigoProducto
	product.Item = req.Item
	product.Cantidad = req.Cantidad
	product.Descripcion = req.Descripcion
	product.Observacion = req.Observacion
	product.Proveedor = req.Proveedor
	product.Ubicacion = req.Ubicacion
	product.Fecha = req.Fecha

	if err := uc.repo.UpdateProduct(product); err != nil {
		return nil, fmt.Errorf("error al actualizar el producto: %v", err)
	}

	resp := &dto.ProductResponse{
		ID:             product.ID,
		CodigoProducto: product.CodigoProducto,
		Item:           product.Item,
		Cantidad:       product.Cantidad,
		Descripcion:    product.Descripcion,
		Observacion:    product.Observacion,
		Proveedor:      product.Proveedor,
		Ubicacion:      product.Ubicacion,
		Fecha:          product.Fecha,
	}

	return resp, nil
}
