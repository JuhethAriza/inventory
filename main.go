package main

import (
	server "github.com/JuhethAriza/inventory/src/infrastructure/server"
	user "github.com/JuhethAriza/inventory/src/modules/User"
)

func main() {
	app := server.ProvidersStore{}
	app.AddModule(user.ModuleProviders())
	app.Up()
}
