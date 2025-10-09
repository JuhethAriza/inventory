package user

import (
	"net/http"

	types "github.com/JuhethAriza/inventory/src/common/types"

	"go.uber.org/fx"
)

func configureModuleRoutes(

	h *types.HandlersStore,
) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Method: http.MethodGet,
				Route:  "/",
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUsers)

}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Invoke(configureModuleRoutes),
	}
}
