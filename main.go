package main

import (
	server "github.com/JuhethAriza/inventory/src/infrastructure/server"
<<<<<<< HEAD
	producto "github.com/JuhethAriza/inventory/src/modules/Producto"
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
	user "github.com/JuhethAriza/inventory/src/modules/User"
)

func main() {
	app := server.ProvidersStore{}
<<<<<<< HEAD
	app.Init()
	app.AddModule(user.ModuleProviders())
	app.AddModule(producto.ModuleProviders())
	app.Up()

=======
	app.AddModule(user.ModuleProviders())
	app.Up()
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
}
