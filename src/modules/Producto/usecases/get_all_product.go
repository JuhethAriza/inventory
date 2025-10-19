package usecases
<<<<<<< HEAD

import (
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	dto "github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
)

type GetAllProducts struct {
	repo *dao.MySQLProductDao
}

func NewGetAllProducts(repo *dao.MySQLProductDao) *GetAllProducts {
	return &GetAllProducts{repo: repo}
}

func (uc *GetAllProducts) Execute() ([]dto.ProductResponse, error) {
	products, err := uc.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var result []dto.ProductResponse
	for _, p := range products {
		result = append(result, dto.ProductResponse{
			ID:             uint(p.ID),
			CodigoProducto: p.CodigoProducto,
			Item:           p.Item,
			Cantidad:       p.Cantidad,
			Categoria:      p.Categoria,
			Estado:         p.Estado,
			Proveedor:      p.Proveedor,
			Fecha:          p.Fecha,
			Ubicacion:      p.Ubicacion,
		})
	}
	return result, nil
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
