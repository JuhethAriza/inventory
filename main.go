package main

import (
	server "github.com/JuhethAriza/inventory/src/infrastructure/server"
	producto "github.com/JuhethAriza/inventory/src/modules/Producto"
	user "github.com/JuhethAriza/inventory/src/modules/User"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(user.ModuleProviders())
	app.AddModule(producto.ModuleProviders())
	app.Up()

}
