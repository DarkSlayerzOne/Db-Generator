package cmd

import (
	appServices "Db-Generator/src/core/application/service"
	portServices "Db-Generator/src/core/ports/service"

	appRouter "Db-Generator/src/core/application/router"
	portRouter "Db-Generator/src/core/ports/router"
	"Db-Generator/src/handlers"
)

var (
	_httpRouter portRouter.MuxRouter = appRouter.MuxRouterImpl()

	_genDb portServices.GenerateMSSQLScriptsServiceInterface = appServices.GenerateMSSQLScriptsServiceInterfaceImpl()

	_handlers handlers.HandlerInterface = handlers.HandlerInterfaceImple(_genDb)
)

func Start() {
	const port string = ":8100"

	_httpRouter.POST("/gen/mssql", _handlers.GenerateDb)

	_httpRouter.SERVE(port)
}
