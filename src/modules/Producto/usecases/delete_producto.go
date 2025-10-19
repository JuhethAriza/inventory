package usecases

import (
	"errors"

	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	repository "github.com/JuhethAriza/inventory/src/modules/Producto/domain/repository"
)

type DeleteProduct struct {
	repo repository.ProductRepository
}

func NewDeleteProduct(repo *dao.MySQLProductDao) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (uc *DeleteProduct) Execute(id int) error {
	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return err
	}
	if product.ID == 0 {
		return errors.New("producto no encontrado")
	}
	return uc.repo.DeleteProduct(id)
}
