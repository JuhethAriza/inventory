package usecases

import (
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	dto "github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
	repository "github.com/JuhethAriza/inventory/src/modules/Producto/domain/repository"
)

type GetProductByID struct {
	repo repository.ProductRepository
}

func NewGetProductByID(repo *dao.MySQLProductDao) *GetProductByID {
	return &GetProductByID{repo: repo}
}

func (uc *GetProductByID) Execute(id int) (*dto.ProductResponse, error) {
	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:             product.ID,
		CodigoProducto: product.CodigoProducto,
		Item:           product.Item,
		Cantidad:       product.Cantidad,
		Descripcion:    product.Descripcion,
		Observacion:    product.Observacion,
		Proveedor:      product.Proveedor,
		Ubicacion:      product.Ubicacion,
		Fecha:          product.Fecha,
	}, nil
}
