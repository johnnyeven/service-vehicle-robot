package routes

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/swagger"
	"github.com/johnnyeven/service-vehicle-robot/routes/v0"
)

var RootRouter = courier.NewRouter(GroupRoot{})

func init() {
	RootRouter.Register(swagger.SwaggerRouter)

	RootRouter.Register(v0.Router)
}

type GroupRoot struct {
	courier.EmptyOperator
}

func (root GroupRoot) Path() string {
	return "/vehicle-robot"
}
