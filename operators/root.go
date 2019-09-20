package operators

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/swagger"
	"github.com/johnnyeven/service-vehicle-robot/operators/nodes"
)

var RootRouter = courier.NewRouter(GroupRoot{})

func init() {
	RootRouter.Register(swagger.SwaggerRouter)
	RootRouter.Register(nodes.Router)
}

type GroupRoot struct {
	courier.EmptyOperator
}

func (root GroupRoot) Path() string {
	return "/vehicle-robot"
}
