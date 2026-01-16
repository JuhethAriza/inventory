package infrastructure

import (
	db "github.com/JuhethAriza/inventory/src/infrastructure/db/adapter"
	"github.com/JuhethAriza/inventory/src/modules/Producto/domain/entities"
	"gorm.io/gorm"
)

type MySQLProductDao struct {
	db *gorm.DB
}

func NewMySQLProductDao(connection *db.DBConnection) *MySQLProductDao {
	return &MySQLProductDao{db: connection.DB}
}

func (dao *MySQLProductDao) CreateProduct(product entities.Product) error {
	query := `
	INSERT INTO products (codigo_producto, item, cantidad, descripcion, observacion, estado, proveedor, fecha, ubicacion)
	VALUES (?, ?, ?, ?, ?,?, ?, ?, ?)
	`
	result := dao.db.Exec(query, product.CodigoProducto, product.Item, product.Cantidad, product.Descripcion, product.Observacion, product.Estado, product.Proveedor, product.Fecha, product.Ubicacion)
	return result.Error
}

func (dao *MySQLProductDao) ExistsByName(item string) (bool, error) {
	var count int64
	err := dao.db.Model(&entities.Product{}).Where("item = ?", item).Count(&count).Error
	return count > 0, err
}

func (dao *MySQLProductDao) DeleteProduct(id int) error {
	result := dao.db.Delete(&entities.Product{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (dao *MySQLProductDao) GetAllProducts() ([]entities.Product, error) {
	var products []entities.Product
	err := dao.db.Find(&products).Error
	return products, err
}

func (dao *MySQLProductDao) UpdateProduct(product entities.Product) error {
	result := dao.db.Model(&entities.Product{}).Where("id = ?", product.ID).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (dao *MySQLProductDao) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product
	err := dao.db.First(&product, "id = ?", id).Error
	return product, err
}
