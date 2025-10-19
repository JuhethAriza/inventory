package usecases
<<<<<<< HEAD

import (
	"errors"

	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	dto "github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
	"github.com/JuhethAriza/inventory/src/modules/Producto/domain/entities"
	repository "github.com/JuhethAriza/inventory/src/modules/Producto/domain/repository"
	"github.com/JuhethAriza/inventory/src/modules/Producto/utils"
)

type CreateProduct struct {
	repo repository.ProductRepository
}

func NewCreateProduct(repo *dao.MySQLProductDao) *CreateProduct {
	return &CreateProduct{
		repo: repo}
}

func (uc *CreateProduct) Execute(request dto.ProductDTO) (*entities.Product, error) {

	if err := utils.ValidateProductDTO(request); err != nil {
		return nil, err
	}

	exists, err := uc.repo.ExistsByName(request.Item)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("ya existe un producto con ese nombre")
	}

	product := entities.Product{
		CodigoProducto: request.CodigoProducto,
		Item:           request.Item,
		Cantidad:       request.Cantidad,
		Categoria:      request.Categoria,
		Estado:         request.Estado,
		Proveedor:      request.Proveedor,
		Fecha:          request.Fecha,
		Ubicacion:      request.Ubicacion,
	}

	if err := uc.repo.CreateProduct(product); err != nil {
		return nil, err
	}

	return &product, nil
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
