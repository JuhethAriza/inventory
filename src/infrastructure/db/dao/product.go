package infrastructure

<<<<<<< HEAD
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
	INSERT INTO products (codigo_producto, item, cantidad, categoria, estado, proveedor, fecha, ubicacion)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result := dao.db.Exec(query, product.CodigoProducto, product.Item, product.Cantidad, product.Categoria, product.Estado, product.Proveedor, product.Fecha, product.Ubicacion)
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
=======
// import (
// 	db "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/adapter"
// 	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/entities"
// 	"gorm.io/gorm"
// )

// type MySQLProductDao struct {
// 	db *gorm.DB
// }

// func NewMySQLProductDao(connection *db.DBConnection) *MySQLProductDao {
// 	return &MySQLProductDao{db: connection.DB}
// }

// func (dao *MySQLProductDao) Create(product *entities.Product) error {
// 	return dao.db.Create(product).Error
// }

// func (dao *MySQLProductDao) ExistsByName(name string) (bool, error) {
// 	var count int64
// 	err := dao.db.Model(&entities.Product{}).Where("name = ?", name).Count(&count).Error
// 	return count > 0, err
// }

// func (dao *MySQLProductDao) GetAll() ([]*entities.Product, error) {
// 	var products []*entities.Product
// 	err := dao.db.Where("status = ?", true).Find(&products).Error
// 	return products, err
// }

// func (dao *MySQLProductDao) GetProductByID(id int) (*entities.Product, error) {
// 	var product entities.Product
// 	err := dao.db.Where("id = ? AND status = ?", id, true).First(&product).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}

// 		return nil, err
// 	}
// 	return &product, nil
// }

// func (dao *MySQLProductDao) UpdateProduct(product *entities.Product) error {
// 	return dao.db.Model(&entities.Product{}).
// 		Where("id = ? AND status = ?", product.ID, true).
// 		Updates(map[string]interface{}{
// 			"name":        product.Name,
// 			"description": product.Description,
// 			"price":       product.Price,
// 			"stock":       product.Stock,
// 		}).Error
// }

// func (dao *MySQLProductDao) DeactivateProduct(id int) error {
// 	return dao.db.Model(&entities.Product{}).
// 		Where("id = ? AND status = ?", id, true).
// 		Update("status", false).Error
// }

// func (dao *MySQLProductDao) ActivateProduct(id int) error {
// 	return dao.db.Model(&entities.Product{}).
// 		Where("id = ? AND status = ?", id, false).
// 		Update("status", true).Error
// }

// func (dao *MySQLProductDao) GetAllDeactivated() ([]*entities.Product, error) {
// 	var products []*entities.Product
// 	err := dao.db.Where("status = ?", false).Find(&products).Error
// 	return products, err
// }
// func (dao *MySQLProductDao) GetProductByIDAnyStatus(id int) (*entities.Product, error) {
// 	var product entities.Product
// 	err := dao.db.Where("id = ?", id).First(&product).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &product, nil
// }

// func (dao *MySQLProductDao) GetLowStock(threshold int) ([]*entities.Product, error) {
// 	var products []*entities.Product
// 	err := dao.db.Where("status = ? AND stock <= ?", true, threshold).Find(&products).Error
// 	return products, err
// }

// func (dao *MySQLProductDao) DeleteProduct(id int) error {
// 	return dao.db.Where("id = ?", id).Delete(&entities.Product{}).Error
// }
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
