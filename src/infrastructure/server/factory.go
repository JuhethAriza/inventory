package infraestructure

import (
	config "github.com/JuhethAriza/inventory/src/common/config"
	result "github.com/JuhethAriza/inventory/src/common/response"
	types "github.com/JuhethAriza/inventory/src/common/types"
	db "github.com/JuhethAriza/inventory/src/infrastructure/db/adapter"

	"go.uber.org/fx"
)

type ProvidersStore struct {
	Providers []fx.Option
}

func (ps *ProvidersStore) Init() {
	ps.Providers = []fx.Option{
		fx.Provide(types.NewHandlersStore),
		fx.Provide(result.NewResult),
		fx.Provide(config.NewConfig),
		fx.Provide(db.NewDBConnection),
	}
}
func (ps *ProvidersStore) AddModule(p []fx.Option) {
	ps.Providers = append(ps.Providers, p...)
}

func (ps *ProvidersStore) Up(lp ...[]fx.Option) {
	ps.Providers = append(ps.Providers, fx.Invoke(NewHttpFiberServer))
	fx.New(ps.Providers...).Run()
}
