package nodes

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-vehicle-robot/modules"
)

func init() {
	Router.Register(courier.NewRouter(RegisterNode{}))
}

// 注册端
type RegisterNode struct {
	httpx.MethodPost
	Body modules.RegisterNodeBody `name:"body" in:"body"`
}

func (req RegisterNode) Path() string {
	return ""
}

func (req RegisterNode) Output(ctx context.Context) (result interface{}, err error) {
	return
}
