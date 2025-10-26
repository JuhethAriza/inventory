package user

import (
	"net/http"

	types "github.com/JuhethAriza/inventory/src/common/types"
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/User/controllers"
	"github.com/JuhethAriza/inventory/src/modules/User/usecases"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlCreateUser *controllers.CreateUsersController,
	ctrlLoginUser *controllers.LoginUserController,
	ctrlGetAllUsers *controllers.GetAllUsersController,
	ctrlGetUserById *controllers.GetUserByIdController,
	h *types.HandlersStore,
) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Method:       http.MethodPost,
				Route:        "/Create",
				Handler:      ctrlCreateUser.Run,
				RequiresAuth: false,
			},
			{
				Method:  http.MethodPost,
				Route:   "/login",
				Handler: ctrlLoginUser.Run,
			},
			{
				Method:  http.MethodGet,
				Route:   "/AllUsers",
				Handler: ctrlGetAllUsers.Run,
			},
			{
				Method:  http.MethodGet,
				Route:   "/:id",
				Handler: ctrlGetUserById.Run,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUsers)

}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLUserDao),
		fx.Provide(controllers.NewCreateUsersController),
		fx.Provide(usecases.NewCreateUsers),
		fx.Provide(controllers.NewLoginUserController),
		fx.Provide(usecases.NewLoginUser),
		fx.Provide(controllers.NewGetAllUsersController),
		fx.Provide(usecases.NewGetAllUsers),
		fx.Provide(controllers.NewGetUserByIdController),
		fx.Provide(usecases.NewGetUserById),
		fx.Invoke(configureModuleRoutes),
	}
}
