package repository

<<<<<<< HEAD
import "github.com/JuhethAriza/inventory/src/modules/Producto/domain/entities"

type ProductRepository interface {
	CreateProduct(product entities.Product) error
	ExistsByName(name string) (bool, error)
	DeleteProduct(id int) error
	GetAllProducts() ([]entities.Product, error)
	UpdateProduct(product entities.Product) error
	GetProductByID(id int) (entities.Product, error)
=======
type ProductRepository interface {
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
}
