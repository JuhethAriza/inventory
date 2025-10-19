package user

import (
	"net/http"

	types "github.com/JuhethAriza/inventory/src/common/types"
<<<<<<< HEAD
	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/JuhethAriza/inventory/src/modules/User/controllers"
	"github.com/JuhethAriza/inventory/src/modules/User/usecases"
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9

	"go.uber.org/fx"
)

func configureModuleRoutes(
<<<<<<< HEAD
	ctrlCreateUser *controllers.CreateUsersController,
	ctrlLoginUser *controllers.LoginUserController,
	ctrlGetAllUsers *controllers.GetAllUsersController,
	ctrlGetUserById *controllers.GetUserByIdController,
=======

>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
	h *types.HandlersStore,
) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
<<<<<<< HEAD
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
=======
				Method: http.MethodGet,
				Route:  "/",
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUsers)

}

func ModuleProviders() []fx.Option {
	return []fx.Option{
<<<<<<< HEAD
		fx.Provide(dao.NewMySQLUserDao),
		fx.Provide(controllers.NewCreateUsersController),
		fx.Provide(usecases.NewCreateUsers),
		fx.Provide(controllers.NewLoginUserController),
		fx.Provide(usecases.NewLoginUser),
		fx.Provide(controllers.NewGetAllUsersController),
		fx.Provide(usecases.NewGetAllUsers),
		fx.Provide(controllers.NewGetUserByIdController),
		fx.Provide(usecases.NewGetUserById),
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
		fx.Invoke(configureModuleRoutes),
	}
}
