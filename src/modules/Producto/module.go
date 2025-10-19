package producto

import (
	"net/http"

	types "github.com/JuhethAriza/inventory/src/common/types"
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/Producto/controllers"
	"github.com/JuhethAriza/inventory/src/modules/Producto/usecases"
	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlCreateProduct *controllers.CreateProductController,
	ctrlGetAllProducts *controllers.GetAllProductsController,
	ctrlUpdateProduct *controllers.UpdateProductController,
	ctrlDeleteProduct *controllers.DeleteProductController,
	ctrlGetProductByID *controllers.GetProductByIDController,
	h *types.HandlersStore,
) {

	handlersModuleProducts := &types.SliceHandlers{
		Prefix: "products",
		Routes: []types.HandlerModule{
			{
				Method:       http.MethodPost,
				Route:        "/Create",
				Handler:      ctrlCreateProduct.Run,
				RequiresAuth: false,
			},
			{
				Method:       http.MethodGet,
				Route:        "/GetAll",
				Handler:      ctrlGetAllProducts.Run,
				RequiresAuth: false,
			},
			{
				Method:       http.MethodPut,
				Route:        "/Update/:id",
				Handler:      ctrlUpdateProduct.Run,
				RequiresAuth: false,
			},
			{
				Method:       http.MethodDelete,
				Route:        "/Delete/:id",
				Handler:      ctrlDeleteProduct.Run,
				RequiresAuth: false,
			},
			{
				Method:       http.MethodGet,
				Route:        "/GetByID/:id",
				Handler:      ctrlGetProductByID.Run,
				RequiresAuth: false,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleProducts)

}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLProductDao),
		fx.Provide(controllers.NewCreateProductController),
		fx.Provide(usecases.NewCreateProduct),
		fx.Provide(controllers.NewGetAllProductsController),
		fx.Provide(usecases.NewGetAllProducts),
		fx.Provide(controllers.NewUpdateProductController),
		fx.Provide(usecases.NewUpdateProduct),
		fx.Provide(controllers.NewDeleteProductController),
		fx.Provide(usecases.NewDeleteProduct),
		fx.Provide(controllers.NewGetProductByIDController),
		fx.Provide(usecases.NewGetProductByID),
		fx.Invoke(configureModuleRoutes),
	}
}
